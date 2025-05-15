package fleetmodel

type FleetHistoryRequest struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}
