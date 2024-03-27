package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/itera-io/openstack-web-client/api/handlers"
	"github.com/itera-io/openstack-web-client/config"
)

func Compute(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewComputeHandler(cfg)

	router.GET("/limits", h.GetComputeLimit)
}
