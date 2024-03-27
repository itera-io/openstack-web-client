package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/itera-io/openstack-web-client/api/dto"
	_ "github.com/itera-io/openstack-web-client/api/helper"
	"github.com/itera-io/openstack-web-client/config"
	"github.com/itera-io/openstack-web-client/services"
)

type ComputeHandler struct {
	service *services.ComputeService
}

func NewComputeHandler(cfg *config.Config) *ComputeHandler {
	service := services.NewComputeService(cfg)
	return &ComputeHandler{service: service}
}

// GetComputeLimit godoc
// @Summary Get Compute Limits
// @Description Get Compute Limits
// @Tags ComputeLimits
// @Accept  json
// @Produce  json
// @Success 200 {object} helper.BaseHttpResponse{result=dto.GetComputeLimitResponse} "GetComputeLimit response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 401 {object} helper.BaseHttpResponse "Unauthorized request"
// @Router /v2/compute/limits [get]
// @Security AuthBearer
func (h *ComputeHandler) GetComputeLimit(c *gin.Context) {
	GetByFilter(c, h.service.GetLimits)
}
