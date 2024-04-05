package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/itera-io/openstack-web-client/api/handlers"
	"github.com/itera-io/openstack-web-client/config"
)

func VolumeType(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewBlockStorageHandler(cfg)

	router.GET("/", h.ListVolumeType)
}
