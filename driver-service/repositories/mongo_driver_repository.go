package repositories

import (
    "context"
    "driver-service/models"
    "errors"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDriverRepository struct {
	Collection *mongo.Collection // büyük harf ile
}


func NewMongoDriverRepository(col *mongo.Collection) DriverRepository {
    return &MongoDriverRepository{
        Collection: col,
    }
}

func (r *MongoDriverRepository) Create(driver *models.Driver) (string, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    driver.CreatedAt = time.Now()
    driver.UpdatedAt = time.Now()

    res, err := r.Collection.InsertOne(ctx, driver)
    if err != nil {
        return "", err
    }

    id := res.InsertedID.(primitive.ObjectID).Hex()
    return id, nil
}

func (r *MongoDriverRepository) Update(id string, driver *models.Driver) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }

    driver.UpdatedAt = time.Now()

    update := bson.M{
        "$set": bson.M{
            "firstName": driver.FirstName,
            "lastName":  driver.LastName,
            "plate":     driver.Plate,
            "taxiType":  driver.TaxiType,
            "carBrand":  driver.CarBrand,
            "carModel":  driver.CarModel,
            "updatedAt": driver.UpdatedAt,
        },
    }

    res, err := r.Collection.UpdateByID(ctx, objID, update)
    if err != nil {
        return err
    }

    if res.MatchedCount == 0 {
        return errors.New("driver not found")
    }

    return nil
}

func (r *MongoDriverRepository) List(page int, pageSize int) ([]models.Driver, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    skip := int64((page - 1) * pageSize)
    limit := int64(pageSize)

    findOptions := options.Find()
    findOptions.SetSkip(skip)
    findOptions.SetLimit(limit)

    cursor, err := r.Collection.Find(ctx, bson.M{}, findOptions)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    drivers := []models.Driver{}
    if err := cursor.All(ctx, &drivers); err != nil {
        return nil, err
    }

    return drivers, nil
}

func (r *MongoDriverRepository) GetByID(id string) (*models.Driver, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }

    var driver models.Driver
    err = r.Collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&driver)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, errors.New("driver not found")
        }
        return nil, err
    }

    return &driver, nil
}
