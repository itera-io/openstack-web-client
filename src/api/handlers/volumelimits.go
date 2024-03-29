package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/itera-io/openstack-web-client/api/dto"
	_ "github.com/itera-io/openstack-web-client/api/helper"
	"github.com/itera-io/openstack-web-client/config"
	"github.com/itera-io/openstack-web-client/services/volume"
)

type VolumeLimitsHandler struct {
	service *volume.Service
}

func NewVolumeLimitsHandler(cfg *config.Config) *VolumeLimitsHandler {
	service := volume.NewService(cfg)
	return &VolumeLimitsHandler{service: service}
}

// GetVolumeLimits godoc
// @Summary Get Volume Limits
// @Description Get Volume Limits
// @Tags VolumeLimits
// @Accept  json
// @Produce  json
// @Success 200 {object} helper.BaseHttpResponse{result=dto.GetVolumeLimitResponse} "GetVolumeLimits response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 401 {object} helper.BaseHttpResponse "Unauthorized request"
// @Router /v3/volume/limits [get]
// @Security AuthBearer
func (h *VolumeLimitsHandler) GetVolumeLimits(c *gin.Context) {
	GetByFilter(c, h.service.GetLimits)
}
