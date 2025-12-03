package redisrepo

import (
    "context"
    "encoding/json"
    "fmt"
    "location-service/models"
    "location-service/config"
    "location-service/calc"
)

type RedisLocationRepository struct {
    rdb *redis.Client
}

func NewRedisLocationRepository() *RedisLocationRepository {
    return &RedisLocationRepository{
        rdb: config.RedisClient,
    }
}

func geoKey(taxiType string) string {
    return fmt.Sprintf("geo:%s", taxiType)
}

func metaKey(driverID string) string {
    return fmt.Sprintf("driver:%s:meta", driverID)
}

// ------------------- UPDATE LOCATION -------------------

func (repo *RedisLocationRepository) SetLocation(driverID string, loc *models.DriverLocation) error {
    ctx := context.Background()

    // GEOADD
    _, err := repo.rdb.GeoAdd(ctx, geoKey(loc.TaxiType), &redis.GeoLocation{
        Name:      driverID,
        Latitude:  loc.Lat,
        Longitude: loc.Lon,
    }).Result()
    if err != nil {
        return err
    }

    // Meta kaydet
    meta, _ := json.Marshal(loc)
    return repo.rdb.Set(ctx, metaKey(driverID), meta, 0).Err()
}

// ------------------- FIND NEARBY -------------------

func (repo *RedisLocationRepository) GetNearby(lat, lon float64, taxiType string) ([]models.NearbyDriver, error) {
    ctx := context.Background()

    // Redis toplanan sürücüler
    results, err := repo.rdb.GeoRadius(ctx, geoKey(taxiType), lon, lat, &redis.GeoRadiusQuery{
        Radius:      5,
        Unit:        "km",
        WithCoord:   true,
        WithDist:    true,
        Count:       20,
        Sort:        "ASC",
    }).Result()
    if err != nil {
        return nil, err
    }

    var list []models.NearbyDriver

    for _, d := range results {
        dist := calc.HaversineDistance(lat, lon, d.Latitude, d.Longitude)

        list = append(list, models.NearbyDriver{
            DriverID: d.Name,
            Lat:      d.Latitude,
            Lon:      d.Longitude,
            Distance: dist,
        })
    }

    return list, nil
}
