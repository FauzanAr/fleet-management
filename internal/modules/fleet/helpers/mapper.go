package fleethelper

import (
	fleetentity "github.com/FauzanAr/fleet-management/internal/modules/fleet/entities"
	fleetmodel "github.com/FauzanAr/fleet-management/internal/modules/fleet/model"
)

func ToFleetResponse(f fleetentity.Fleet) fleetmodel.FleetResponse {
	return fleetmodel.FleetResponse{
		VehicleId: f.VehicleId,
		Latitude:  f.Latitude,
		Logitude:  f.Longitude,
		Timestamp: f.Timestamp,
	}
}

func ToFleetResponseList(fleets []fleetentity.Fleet) []fleetmodel.FleetResponse {
	responses := make([]fleetmodel.FleetResponse, 0, len(fleets))
	for _, f := range fleets {
		responses = append(responses, ToFleetResponse(f))
	}

	return responses
}
