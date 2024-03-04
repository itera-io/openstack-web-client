package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/itera-io/openstack-web-client/api/handlers"
	"github.com/itera-io/openstack-web-client/config"
)

func Flavor(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewFlavorsHandler(cfg)

	router.GET("/", h.ListFlavors)
}
