package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/itera-io/openstack-web-client/api/handlers"
	"github.com/itera-io/openstack-web-client/config"
)

func Role(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewRolesHandler(cfg)

	router.GET("/", h.ListRoles)
	router.PUT("/:id/", h.AssignRole)
}
