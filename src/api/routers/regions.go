package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/itera-io/openstack-web-client/api/handlers"
)

func Region(router *gin.RouterGroup) {
	h := handlers.NewRegionsHandler()

	router.GET("/", h.ListRegions)
}
