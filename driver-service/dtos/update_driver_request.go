package dtos

type UpdateDriverRequest struct {
    FirstName string `json:"firstName,omitempty"`
    LastName  string `json:"lastName,omitempty"`
    Plate     string `json:"plate,omitempty"`
    TaxiType  string `json:"taxiType,omitempty"`
    CarBrand  string `json:"carBrand,omitempty"`
    CarModel  string `json:"carModel,omitempty"`
}
