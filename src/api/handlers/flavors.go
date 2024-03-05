package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itera-io/openstack-web-client/api/dto"
	"github.com/itera-io/openstack-web-client/api/helper"
	"github.com/itera-io/openstack-web-client/config"
	"github.com/itera-io/openstack-web-client/constants"
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
// @Param Request body dto.ListFlavorRequest true "ListFlavorRequest"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.ListFlavorResponse} "ListFlavor response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 401 {object} helper.BaseHttpResponse "Unauthorized request"
// @Router /v2/flavors [get]
// @Security AuthBearer
func (h *FlavorsHandler) ListFlavors(c *gin.Context) {
	var t, _ = c.Get(constants.TokenKey)
	var u, _ = c.Get(constants.AuthUrlKey)
	authUtils := &dto.AuthUtils{Token: t.(string), BaseUrl: u.(string)}
	res, err := h.service.ListFlavors(new(dto.ListFlavorRequest), authUtils)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, helper.Success))
}

// GetFlavor godoc
// @Summary Get Flavor
// @Description Get Flavor
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
	id := c.Params.ByName("id")
	if id == "" {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponse(nil, false, helper.ValidationError))
		return
	}
	var t, _ = c.Get(constants.TokenKey)
	var u, _ = c.Get(constants.AuthUrlKey)
	authUtils := &dto.AuthUtils{Token: t.(string), BaseUrl: u.(string)}
	res, err := h.service.GetFlavor(id, authUtils)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, helper.Success))
}
