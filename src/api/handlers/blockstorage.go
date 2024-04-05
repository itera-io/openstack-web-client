package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/itera-io/openstack-web-client/api/dto"
	_ "github.com/itera-io/openstack-web-client/api/helper"
	"github.com/itera-io/openstack-web-client/config"
	"github.com/itera-io/openstack-web-client/services/blockstorage"
)

type BlockStorageHandler struct {
	service *blockstorage.Service
}

func NewBlockStorageHandler(cfg *config.Config) *BlockStorageHandler {
	service := blockstorage.NewService(cfg)
	return &BlockStorageHandler{service: service}
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
func (h *BlockStorageHandler) GetVolumeLimits(c *gin.Context) {
	GetByFilter(c, h.service.GetLimits)
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
func (h *BlockStorageHandler) ListVolumeType(c *gin.Context) {
	GetByFilter(c, h.service.ListTypes)
}
