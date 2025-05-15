package fleetmodel

type FleetResponse struct {
	VehicleId string `json:"vehicle_id"`
	Latitude  int64  `json:"latitude"`
	Logitude  int64  `json:"longitude"`
	Timestamp int64  `json:"timestamp"`
}
