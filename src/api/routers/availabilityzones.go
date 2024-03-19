package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/itera-io/openstack-web-client/api/handlers"
	"github.com/itera-io/openstack-web-client/config"
)

func AvailabilityZone(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewAvailabilityZonesHandler(cfg)

	router.GET("/", h.List)
}
