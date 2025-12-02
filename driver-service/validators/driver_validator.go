package validator

import (
	"driver-service/dtos"
	"regexp"
	"strings"
)

type DriverValidator struct{}

var allowedTaxiTypes = map[string]bool{
	"sari":    true,
	"siyah":   true,
	"premium": true,
}

var allowedCarModels = map[string]bool{
	"long":     true,
	"short":    true,
	"standard": true,
}

func NewDriverValidator() *DriverValidator {
	return &DriverValidator{}
}

// ValidateCreateDriver artık ValidationResponse döndürüyor
func (v *DriverValidator) ValidateCreateDriver(req *dtos.CreateDriverRequest) *dtos.ValidationResponse {
	var validationErrors []dtos.ValidationError

	if strings.TrimSpace(req.FirstName) == "" {
		validationErrors = append(validationErrors, dtos.ValidationError{"firstName", "first name is required"})
	}
	if strings.TrimSpace(req.LastName) == "" {
		validationErrors = append(validationErrors, dtos.ValidationError{"lastName", "last name is required"})
	}
	if !isValidPlate(req.Plate) {
		validationErrors = append(validationErrors, dtos.ValidationError{"plate", "invalid plate format"})
	}
	if !allowedTaxiTypes[strings.ToLower(req.TaxiType)] {
		validationErrors = append(validationErrors, dtos.ValidationError{"taxiType", "invalid taxi type, allowed: sari, siyah, premium"})
	}
	if req.CarBrand == "" {
		validationErrors = append(validationErrors, dtos.ValidationError{"carBrand", "car brand is required"})
	}
	if !allowedCarModels[strings.ToLower(req.CarModel)] {
		validationErrors = append(validationErrors, dtos.ValidationError{"carModel", "invalid car model, allowed: long, short, standard"})
	}
	if strings.TrimSpace(req.UserId) == "" {
		validationErrors = append(validationErrors, dtos.ValidationError{"userId", "user id is required"})
	}

	if len(validationErrors) > 0 {
		return &dtos.ValidationResponse{Errors: validationErrors}
	}
	return nil
}

// UpdateDriver versiyonu aynı şekilde
func (v *DriverValidator) ValidateUpdateDriver(req *dtos.UpdateDriverRequest) *dtos.ValidationResponse {
	var validationErrors []dtos.ValidationError

	if strings.TrimSpace(req.FirstName) == "" {
		validationErrors = append(validationErrors, dtos.ValidationError{"firstName", "first name is required"})
	}
	if strings.TrimSpace(req.LastName) == "" {
		validationErrors = append(validationErrors, dtos.ValidationError{"lastName", "last name is required"})
	}
	if !isValidPlate(req.Plate) {
		validationErrors = append(validationErrors, dtos.ValidationError{"plate", "invalid plate format"})
	}
	if !allowedTaxiTypes[strings.ToLower(req.TaxiType)] {
		validationErrors = append(validationErrors, dtos.ValidationError{"taxiType", "invalid taxi type, allowed: sari, siyah, premium"})
	}
	if req.CarBrand == "" {
		validationErrors = append(validationErrors, dtos.ValidationError{"carBrand", "car brand is required"})
	}
	if !allowedCarModels[strings.ToLower(req.CarModel)] {
		validationErrors = append(validationErrors, dtos.ValidationError{"carModel", "invalid car model, allowed: long, short, standard"})
	}

	if len(validationErrors) > 0 {
		return &dtos.ValidationResponse{Errors: validationErrors}
	}
	return nil
}

func isValidPlate(plate string) bool {
	re := regexp.MustCompile(`^[A-Z0-9-]{2,10}$`)
	return re.MatchString(strings.ToUpper(plate))
}
