package fleethandler

import (
	"net/http"
	"strconv"

	"github.com/FauzanAr/fleet-management/internal/modules/fleet"
	fleetmodel "github.com/FauzanAr/fleet-management/internal/modules/fleet/model"
	"github.com/FauzanAr/fleet-management/internal/pkg/logger"
	"github.com/FauzanAr/fleet-management/internal/pkg/wrapper"
	"github.com/gin-gonic/gin"
)

type FleetHandler struct {
	log logger.Logger
	fu  fleet.Usecase
}

func NewFleetHandlers(log logger.Logger, fu fleet.Usecase) *FleetHandler {
	return &FleetHandler{
		log: log,
		fu:  fu,
	}
}

func (h *FleetHandler) LastLocation(c *gin.Context) {
	var request fleetmodel.FleetLastLocationRequest
	ctx := c.Request.Context()

	vehicleId := c.Param("vehicle_id")
	if vehicleId == "" {
		wrapper.SendErrorResponse(c, wrapper.ValidationError("vehicle_id path param is required"), nil, http.StatusBadRequest)
		return
	}

	request.VehicleId = vehicleId

	response, err := h.fu.GetFleet(ctx, request)
	if err != nil {
		wrapper.SendErrorResponse(c, err, nil, http.StatusBadRequest)
		return
	}

	wrapper.SendSuccessResponse(c, "Success", response, http.StatusOK)
	return
}

func (h *FleetHandler) History(c *gin.Context) {
	var request fleetmodel.FleetHistoryRequest
	ctx := c.Request.Context()

	startStr := c.Query("start")
	endStr := c.Query("end")
	vehicleId := c.Param("vehicle_id")

	if vehicleId == "" {
		wrapper.SendErrorResponse(c, wrapper.ValidationError("vehicle_id path param is required"), nil, http.StatusBadRequest)
		return
	}

	if startStr == "" || endStr == "" {
		wrapper.SendErrorResponse(c, wrapper.ValidationError("start or end query params are required"), nil, http.StatusBadRequest)
		return
	}

	start, err := strconv.ParseInt(startStr, 10, 64)
	if err != nil {
		wrapper.SendErrorResponse(c, wrapper.ValidationError("start query param must be a number"), nil, http.StatusBadRequest)
		return
	}

	end, err := strconv.ParseInt(endStr, 10, 64)
	if err != nil {
		wrapper.SendErrorResponse(c, wrapper.ValidationError("end query param must be a number"), nil, http.StatusBadRequest)
		return
	}

	if end < start {
		wrapper.SendErrorResponse(c, wrapper.ValidationError("end query param must be greater than start"), nil, http.StatusBadRequest)
		return
	}

	request.Start = start
	request.End = end
	request.VehicleId = vehicleId

	response, err := h.fu.GetFleetHistory(ctx, request)
	if err != nil {
		wrapper.SendErrorResponse(c, err, nil, http.StatusBadRequest)
		return
	}

	wrapper.SendSuccessResponse(c, "Success", response, http.StatusOK)
	return
}
