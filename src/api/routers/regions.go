package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/itera-io/openstack-web-client/api/handlers"
	"github.com/itera-io/openstack-web-client/config"
)

func Region(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewRegionsHandler(cfg)

	router.GET("/", h.ListRegions)
}
