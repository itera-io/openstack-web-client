package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/itera-io/openstack-web-client/api/handlers"
	"github.com/itera-io/openstack-web-client/config"
)

func ApiVersion(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewApiVersionsHandler(cfg)

	router.GET("/", h.List)
}
