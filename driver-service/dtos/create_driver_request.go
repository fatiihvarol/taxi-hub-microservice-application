package dtos

type CreateDriverRequest struct {
    FirstName string `json:"firstName"`
    LastName  string `json:"lastName"`
    Plate     string `json:"plate"`
    TaxiType  string `json:"taxiType"`
    CarBrand  string `json:"carBrand"`
    CarModel  string `json:"carModel"`
    UserId    string `json:"userId"`
}
