package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/itera-io/openstack-web-client/api/dto"
	_ "github.com/itera-io/openstack-web-client/api/helper"
	"github.com/itera-io/openstack-web-client/config"
	"github.com/itera-io/openstack-web-client/services"
)

type VolumeTypesHandler struct {
	service *services.VolumeTypeService
}

func NewVolumeTypesHandler(cfg *config.Config) *VolumeTypesHandler {
	service := services.NewVolumeTypeService(cfg)
	return &VolumeTypesHandler{service: service}
}

// ListVolumeType godoc
// @Summary List VolumeType
// @Description List VolumeType
// @Tags VolumeTypes
// @Accept  json
// @Produce  json
// @Param Request body dto.ListVolumeTypeRequest true "ListVolumeType Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.ListVolumeTypeResponse} "ListVolumeType response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 401 {object} helper.BaseHttpResponse "Unauthorized request"
// @Router /v3/volumetypes [get]
// @Security AuthBearer
func (h *VolumeTypesHandler) List(c *gin.Context) {
	GetByFilter(c, h.service.List)
}
