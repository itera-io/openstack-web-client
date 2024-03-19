package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/itera-io/openstack-web-client/api/dto"
	_ "github.com/itera-io/openstack-web-client/api/helper"
	"github.com/itera-io/openstack-web-client/config"
	"github.com/itera-io/openstack-web-client/services"
)

type ApiVersionsHandler struct {
	service *services.ApiVersionService
}

func NewApiVersionsHandler(cfg *config.Config) *ApiVersionsHandler {
	service := services.NewApiVersionService(cfg)
	return &ApiVersionsHandler{service: service}
}

// ListApiVersion godoc
// @Summary List ApiVersion
// @Description List ApiVersion
// @Tags ApiVersions
// @Accept  json
// @Produce  json
// @Param Request body dto.ListApiVersionRequest true "ListApiVersion Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.ListApiVersionResponse} "ListApiVersion response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 401 {object} helper.BaseHttpResponse "Unauthorized request"
// @Router /v2/apiversions [get]
// @Security AuthBearer
func (h *ApiVersionsHandler) List(c *gin.Context) {
	GetByFilter(c, h.service.ListApiVersions)
}
