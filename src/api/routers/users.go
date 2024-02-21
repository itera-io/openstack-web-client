package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/itera-io/openstack-web-client/api/handlers"
)

func User(router *gin.RouterGroup) {
	h := handlers.NewUsersHandler()

	router.POST("/validate", h.ValidateUser)
}
