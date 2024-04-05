package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/itera-io/openstack-web-client/api/dto"
	_ "github.com/itera-io/openstack-web-client/api/helper"
	"github.com/itera-io/openstack-web-client/config"
	"github.com/itera-io/openstack-web-client/services/compute"
)

type AvailabilityZonesHandler struct {
	service *compute.Service
}

func NewAvailabilityZonesHandler(cfg *config.Config) *AvailabilityZonesHandler {
	service := compute.NewService(cfg)
	return &AvailabilityZonesHandler{service: service}
}

// ListAvailabilityZone godoc
// @Summary List AvailabilityZone
// @Description List AvailabilityZone
// @Tags AvailabilityZones
// @Accept  json
// @Produce  json
// @Param Request body dto.ListAvailabilityZoneRequest true "ListAvailabilityZone Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.ListAvailabilityZoneResponse} "ListAvailabilityZone response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 401 {object} helper.BaseHttpResponse "Unauthorized request"
// @Router /v2/availabilityzones [get]
// @Security AuthBearer
func (h *AvailabilityZonesHandler) List(c *gin.Context) {
	GetByFilter(c, h.service.List)
}
