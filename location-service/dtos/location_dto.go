package dtos

type UpdateLocationRequest struct {
    DriverID  string  `json:"driverId"` 
    Lat       float64 `json:"lat"`
    Lon       float64 `json:"lon"`
    TaxiType  string  `json:"taksiType"`
}
