package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/itera-io/openstack-web-client/api/handlers"
	"github.com/itera-io/openstack-web-client/config"
)

func VolumeLimits(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewVolumeLimitsHandler(cfg)

	router.GET("/limits", h.GetVolumeLimits)
}
