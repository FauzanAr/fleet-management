package fleethandler

import "github.com/gin-gonic/gin"

func (h *FleetHandler) FleetRoutes(router *gin.RouterGroup) {
	router.GET("/:vehicle_id/location", h.LastLocation)
	router.GET("/:vehicle_id/history", h.History)
}