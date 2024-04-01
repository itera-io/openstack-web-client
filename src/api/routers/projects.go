package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/itera-io/openstack-web-client/api/handlers"
	"github.com/itera-io/openstack-web-client/config"
)

func Project(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewProjectsHandler(cfg)

	router.GET("/", h.ListProjects)
	router.POST("/", h.CreateProject)
}
