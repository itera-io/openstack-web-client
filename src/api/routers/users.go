package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/itera-io/openstack-web-client/api/handlers"
	"github.com/itera-io/openstack-web-client/config"
)

func User(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewUsersHandler(cfg)

	router.POST("/validate", h.Validate)
	router.POST("/auth", h.Authenticate)
}

func AuthenticatedUser(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewUsersHandler(cfg)

	router.GET("/:id/projects", h.ListUserProjects)
	router.POST("/", h.CreateUser)
}
