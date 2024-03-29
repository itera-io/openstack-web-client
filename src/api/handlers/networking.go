package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/itera-io/openstack-web-client/api/dto"
	_ "github.com/itera-io/openstack-web-client/api/helper"
	"github.com/itera-io/openstack-web-client/config"
	"github.com/itera-io/openstack-web-client/services/networking"
)

type NetworkingHandler struct {
	service *networking.Service
}

func NewNetworkingHandler(cfg *config.Config) *NetworkingHandler {
	service := networking.NewService(cfg)
	return &NetworkingHandler{service: service}
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
func (h *NetworkingHandler) GetQuotas(c *gin.Context) {
	GetByFilter(c, h.service.GetQuotas)
}

// ListNetworkSubnets godoc
// @Summary List Network Subnets
// @Description List Network Subnets
// @Tags NetworkSubnets
// @Accept  json
// @Produce  json
// @Param Request body dto.ListNetworkSubnetRequest true "ListNetworkSubnet Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.ListNetworkSubnetResponse} "ListNetworkSubnet response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 401 {object} helper.BaseHttpResponse "Unauthorized request"
// @Router /v2/networking/subnets [get]
// @Security AuthBearer
func (h *NetworkingHandler) GetSubnets(c *gin.Context) {
	GetByFilter(c, h.service.GetSubnets)
}

// ListNetworks godoc
// @Summary List Networks
// @Description List Networks
// @Tags Networks
// @Accept  json
// @Produce  json
// @Param Request body dto.ListNetworkRequest true "ListNetwork Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.ListNetworkResponse} "ListNetwork response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 401 {object} helper.BaseHttpResponse "Unauthorized request"
// @Router /v2/networking/networks [get]
// @Security AuthBearer
func (h *NetworkingHandler) GetNetworks(c *gin.Context) {
	GetByFilter(c, h.service.GetNetworks)
}
