package services

import (
	"driver-service/models"
	"driver-service/repositories"
)

type DriverService struct {
	Repo repositories.DriverRepository
}

// Repo gönderiliyor
func NewDriverService(repo repositories.DriverRepository) *DriverService {
	return &DriverService{Repo: repo}
}

func (s *DriverService) CreateDriver(driver *models.Driver) (string, error) {
    id, err := s.Repo.Create(driver)
    if err != nil {
        return "", err
    }
    return id, nil
}


func (s *DriverService) UpdateDriver(id string, driver *models.Driver) (*models.Driver, error) {
	err := s.Repo.Update(id, driver)
	if err != nil {
		return nil, err
	}
	driver.ID = id
	return driver, nil
}

func (s *DriverService) ListDrivers(page, pageSize int) ([]models.Driver, error) {
	return s.Repo.List(page, pageSize)
}

func (s *DriverService) GetDriverByID(id string) (*models.Driver, error) {
	return s.Repo.GetByID(id)
}

// Nearby driver fonksiyonu örnek (placeholder)
func (s *DriverService) GetNearby(lat, lon float64, taxiType string) ([]models.Driver, error) {
	// MongoDB query veya business logic eklenebilir
	return s.Repo.List(1, 10) // basit örnek
}
