package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/itera-io/openstack-web-client/api/handlers"
	"github.com/itera-io/openstack-web-client/config"
)

func Server(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewServersHandler(cfg)

	router.GET("/", h.ListServers)
	router.GET("/:id/", h.GetServer)
	router.PUT("/:id/reboot/", h.RebootServer)
	router.PUT("/:id/start/", h.StartServer)
	router.PUT("/:id/stop/", h.StopServer)
	router.PUT("/:id/shelve/", h.ShelveServer)
	router.PUT("/:id/unshelve/", h.UnshelveServer)
}
