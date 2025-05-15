package fleetentity

type Fleet struct {
	ID        uint64  `gorm:"primaryKey"`
	VehicleId string  `gorm:"column:vehicle_id"`
	Latitude  float64 `gorm:"column:latitude"`
	Longitude float64 `gorm:"column:longitude"`
	Timestamp int64   `gorm:"column:timestamp"`
}

type FleetHistoryQuery struct {
	Start     int64
	End       int64
	VehicleId string
}

func (Fleet) TableName() string {
	return "vehicle_locations"
}
