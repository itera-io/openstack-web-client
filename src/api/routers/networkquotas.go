package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/itera-io/openstack-web-client/api/handlers"
	"github.com/itera-io/openstack-web-client/config"
)

func NetworkQuotas(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewNetworkQuotasHandler(cfg)

	router.GET("/quotas", h.Get)
}
