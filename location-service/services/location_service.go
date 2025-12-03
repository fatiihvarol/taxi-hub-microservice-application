package service

import (
    "location-service/models"
    redisrepo "location-service/infrastructure/redis"
)

type LocationService struct {
    repo *redisrepo.RedisLocationRepository
}

func NewLocationService(r *redisrepo.RedisLocationRepository) *LocationService {
    return &LocationService{repo: r}
}

func (s *LocationService) Update(driverID string, loc *models.DriverLocation) error {
    return s.repo.SetLocation(driverID, loc)
}

func (s *LocationService) Nearby(lat, lon float64, taxiType string) ([]models.NearbyDriver, error) {
    return s.repo.GetNearby(lat, lon, taxiType)
}
