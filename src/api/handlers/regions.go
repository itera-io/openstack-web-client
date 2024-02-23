package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itera-io/openstack-web-client/api/dto"
	"github.com/itera-io/openstack-web-client/api/helper"
	"github.com/itera-io/openstack-web-client/constants"
	"github.com/itera-io/openstack-web-client/services"
)

type RegionsHandler struct {
	service *services.RegionService
}

func NewRegionsHandler() *RegionsHandler {
	service := services.NewRegionService()
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
	req := new(dto.ListRegionRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}
	var t, _ = c.Get(constants.TokenKey)
	var u, _ = c.Get(constants.BaseUrlKey)
	authUtils := &dto.AuthUtils{Token: t.(string), BaseUrl: u.(string)}
	res, err := h.service.ListRegions(req, authUtils)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, helper.Success))
}
