package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/itera-io/openstack-web-client/api/handlers"
	"github.com/itera-io/openstack-web-client/config"
)

func Image(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewImagesHandler(cfg)

	router.GET("/", h.ListImages)
	router.GET("/:id/", h.GetImage)
}
