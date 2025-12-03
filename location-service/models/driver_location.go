package models

type DriverLocation struct {
    Lat      float64 `json:"lat"`
    Lon      float64 `json:"lon"`
    TaxiType string  `json:"taksiType"`
    UpdatedAt int64  `json:"updatedAt"`
}

type NearbyDriver struct {
    DriverID string  `json:"driverId"`
    Lat      float64 `json:"lat"`
    Lon      float64 `json:"lon"`
    TaxiType string  `json:"taksiType"`
    Distance float64 `json:"distance"`
}
