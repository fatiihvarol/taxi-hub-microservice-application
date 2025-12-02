package services

import (
	"driver-service/dtos"
	"driver-service/models"
	"driver-service/repositories"
	"time"
)

type DriverService struct {
	Repo repositories.DriverRepository
}

func NewDriverService(repo repositories.DriverRepository) *DriverService {
	return &DriverService{Repo: repo}
}

// CreateDriver artık DTO alıyor
func (s *DriverService) CreateDriver(req *dtos.CreateDriverRequest) (*dtos.CreateDriverResponse, error) {
	driver := &models.Driver{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Plate:     req.Plate,
		TaxiType:  req.TaxiType,
		CarBrand:  req.CarBrand,
		CarModel:  req.CarModel,
		UserId:    req.UserId,
	}

	id, err := s.Repo.Create(driver)
	if err != nil {
		return nil, err
	}

return &dtos.CreateDriverResponse{
    ID: id,
}, nil
}

// UpdateDriver artık DTO alıyor
func (s *DriverService) UpdateDriver(id string, req *dtos.UpdateDriverRequest) (*dtos.UpdateDriverResponse, error) {
	driver := &models.Driver{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Plate:     req.Plate,
		TaxiType:  req.TaxiType,
		CarBrand:  req.CarBrand,
		CarModel:  req.CarModel,
	}

	err := s.Repo.Update(id, driver)
	if err != nil {
		return nil, err
	}
	driver.ID = id

return &dtos.UpdateDriverResponse{
    ID:        driver.ID,
    UpdatedAt: time.Now().Format(time.RFC3339),
}, nil
}

func (s *DriverService) ListDrivers(page, pageSize int) ([]models.Driver, error) {
	return s.Repo.List(page, pageSize)
}

func (s *DriverService) GetDriverByID(id string) (*models.Driver, error) {
	return s.Repo.GetByID(id)
}

func (s *DriverService) GetNearby(lat, lon float64, taxiType string) ([]models.Driver, error) {
	return s.Repo.List(1, 10) // placeholder
}
