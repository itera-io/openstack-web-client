package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/itera-io/openstack-web-client/api/dto"
	_ "github.com/itera-io/openstack-web-client/api/helper"
	"github.com/itera-io/openstack-web-client/config"
	"github.com/itera-io/openstack-web-client/services"
)

type NetworkQuotasHandler struct {
	service *services.NetworkQuotaService
}

func NewNetworkQuotasHandler(cfg *config.Config) *NetworkQuotasHandler {
	service := services.NewNetworkQuotaService(cfg)
	return &NetworkQuotasHandler{service: service}
}

// GetNetworkQuotas godoc
// @Summary Get Network Quotas
// @Description Get Network Quotas
// @Tags NetworkQuotas
// @Accept  json
// @Produce  json
// @Param Request body dto.GetNetworkQuotaRequest true "GetNetworkQuota Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.GetNetworkQuotaResponse} "GetNetworkQuota response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 401 {object} helper.BaseHttpResponse "Unauthorized request"
// @Router /v2/networking/quotas [get]
// @Security AuthBearer
func (h *NetworkQuotasHandler) Get(c *gin.Context) {
	GetByFilter(c, h.service.Get)
}
