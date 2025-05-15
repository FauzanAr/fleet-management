package fleetmodel

type FleetResponse struct {
	VehicleId string `json:"vehicle_id"`
	Latitude  float64  `json:"latitude"`
	Logitude  float64  `json:"longitude"`
	Timestamp int64  `json:"timestamp"`
}
