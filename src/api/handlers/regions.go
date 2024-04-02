package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/itera-io/openstack-web-client/api/dto"
	_ "github.com/itera-io/openstack-web-client/api/helper"
	"github.com/itera-io/openstack-web-client/config"
	v3 "github.com/itera-io/openstack-web-client/services/identity/v3"
)

type RegionsHandler struct {
	service *v3.Service
}

func NewRegionsHandler(cfg *config.Config) *RegionsHandler {
	service := v3.NewService(cfg)
	return &RegionsHandler{service: service}
}

// ListRegion godoc
// @Summary List Region
// @Description List Region
// @Tags Regions
// @Accept  json
// @Produce  json
// @Param Request body dto.ListRegionRequest true "ListRegionRequest"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.ListRegionResponse} "ListRegion response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 401 {object} helper.BaseHttpResponse "Unauthorized request"
// @Router /v3/regions [get]
// @Security AuthBearer
func (h *RegionsHandler) ListRegions(c *gin.Context) {
	GetByFilter(c, h.service.ListRegions)
}
