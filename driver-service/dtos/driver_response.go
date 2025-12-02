package dtos

type DriverResponse struct {
    ID        string `json:"id"`
    FirstName string `json:"firstName"`
    LastName  string `json:"lastName"`
    Plate     string `json:"plate"`
    TaxiType  string `json:"taxiType"`
    CarBrand  string `json:"carBrand"`
    CarModel  string `json:"carModel"`
    CreatedAt string `json:"createdAt"`
    UpdatedAt string `json:"updatedAt"`
}
