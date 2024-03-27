package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/itera-io/openstack-web-client/api/handlers"
	"github.com/itera-io/openstack-web-client/config"
)

func ComputeLimits(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewComputeLimitsHandler(cfg)

	router.GET("/limits", h.GetComputeLimits)
}
