package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/itera-io/openstack-web-client/api/dto"
	_ "github.com/itera-io/openstack-web-client/api/helper"
	"github.com/itera-io/openstack-web-client/config"
	"github.com/itera-io/openstack-web-client/services/compute"
)

type ServersHandler struct {
	service *compute.Service
}

func NewServersHandler(cfg *config.Config) *ServersHandler {
	service := compute.NewService(cfg)
	return &ServersHandler{service: service}
}

// GetServer godoc
// @Summary Get Server
// @Description Get Server by ID
// @Tags Servers
// @Accept  json
// @Produce  json
// @Param id path string true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.GetServerResponse} "GetServer response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 401 {object} helper.BaseHttpResponse "Unauthorized request"
// @Router /v2/servers/{id} [get]
// @Security AuthBearer
func (h *ServersHandler) GetServer(c *gin.Context) {
	GetById(c, h.service.GetServer)
}
