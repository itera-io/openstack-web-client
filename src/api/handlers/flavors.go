package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/itera-io/openstack-web-client/api/dto"
	_ "github.com/itera-io/openstack-web-client/api/helper"
	"github.com/itera-io/openstack-web-client/config"
	"github.com/itera-io/openstack-web-client/services"
)

type FlavorsHandler struct {
	service *services.FlavorService
}

func NewFlavorsHandler(cfg *config.Config) *FlavorsHandler {
	service := services.NewFlavorService(cfg)
	return &FlavorsHandler{service: service}
}

// ListFlavor godoc
// @Summary List Flavor
// @Description List Flavor
// @Tags Flavors
// @Accept  json
// @Produce  json
// @Param Request body dto.ListFlavorRequest true "ListFlavor Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.ListFlavorResponse} "ListFlavor response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 401 {object} helper.BaseHttpResponse "Unauthorized request"
// @Router /v2/flavors [get]
// @Security AuthBearer
func (h *FlavorsHandler) ListFlavors(c *gin.Context) {
	GetByFilter(c, h.service.ListFlavors)
}

// GetFlavor godoc
// @Summary Get Flavor
// @Description Get Flavor by ID
// @Tags Flavors
// @Accept  json
// @Produce  json
// @Param id path string true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.GetFlavorResponse} "GetFlavor response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 401 {object} helper.BaseHttpResponse "Unauthorized request"
// @Router /v2/flavors/{id} [get]
// @Security AuthBearer
func (h *FlavorsHandler) GetFlavor(c *gin.Context) {
	GetById(c, h.service.GetFlavor)
}
