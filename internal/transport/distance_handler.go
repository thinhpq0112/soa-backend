package transport

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/thinhpq0112/soa-backend/internal/service"
	"net/http"
	"strings"
)

type DistanceHandler struct {
	distanceService *service.DistanceService
}

func NewDistanceHandler(distanceService *service.DistanceService) *DistanceHandler {
	return &DistanceHandler{distanceService: distanceService}
}

// @Accept json
// @Produce json
// @Param city query string true "City name"
// @Success 200 {object} model.DistanceResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/distance [get]
func (h *DistanceHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/distance", h.CalculateDistanceHandler)
}

func (h *DistanceHandler) CalculateDistanceHandler(c *gin.Context) {
	ip := c.ClientIP()
	if ip == "" {
		ip = strings.Split(c.Request.RemoteAddr, ":")[0]
	}

	city := c.Query("city")
	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cicty is required"})
		return
	}

	distance, err := h.distanceService.CalculateDistance(context.Background(), ip, city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"distance_km": distance})
}
