package repositories

import (
	"driver-service/models"
)

type DriverRepository interface {
	Create(driver *models.Driver) (string, error)
	Update(id string, driver *models.Driver) error
	List(page int, pageSize int) ([]models.Driver, error)
	GetByID(id string) (*models.Driver, error)
}
