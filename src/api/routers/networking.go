package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/itera-io/openstack-web-client/api/handlers"
	"github.com/itera-io/openstack-web-client/config"
)

func Networking(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewNetworkingHandler(cfg)

	router.GET("/quotas", h.GetQuotas)
	router.GET("/subnets", h.GetSubnets)
	router.GET("/networks", h.GetNetworks)
}
