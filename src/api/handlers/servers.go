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

// ListServer godoc
// @Summary List Server
// @Description List Server
// @Tags Servers
// @Accept  json
// @Produce  json
// @Param Request body dto.ListServerRequest true "ListServer Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.ListServerResponse} "ListServer response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 401 {object} helper.BaseHttpResponse "Unauthorized request"
// @Router /v2/servers [get]
// @Security AuthBearer
func (h *ServersHandler) ListServers(c *gin.Context) {
	GetByFilter(c, h.service.ListServers)
}

// RebootServer godoc
// @Summary Reboot Server
// @Description Reboot Server
// @Tags Servers
// @Accept  json
// @Produce  json
// @Param id path string true "Id"
// @Param Request body dto.RebootServerRequest true "RebootServerRequest"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.RebootServerResponse} "RebootServer response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 401 {object} helper.BaseHttpResponse "Unauthorized request"
// @Router /v2/servers/{id} [put]
// @Security AuthBearer
func (h *ServersHandler) RebootServer(c *gin.Context) {
	Update(c, h.service.RebootServer)
}

// StartServer godoc
// @Summary Start Server
// @Description Start Server
// @Tags Servers
// @Accept  json
// @Produce  json
// @Param id path string true "Id"
// @Success 200 {object} helper.BaseHttpResponse "StartServer response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 401 {object} helper.BaseHttpResponse "Unauthorized request"
// @Router /v2/servers/{id}/start [put]
// @Security AuthBearer
func (h *ServersHandler) StartServer(c *gin.Context) {
	UpdateWithoutBody(c, h.service.StartServer)
}

// StopServer godoc
// @Summary Stop Server
// @Description Stop Server
// @Tags Servers
// @Accept  json
// @Produce  json
// @Param id path string true "Id"
// @Success 200 {object} helper.BaseHttpResponse "StopServer response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 401 {object} helper.BaseHttpResponse "Unauthorized request"
// @Router /v2/servers/{id}/stop [put]
// @Security AuthBearer
func (h *ServersHandler) StopServer(c *gin.Context) {
	UpdateWithoutBody(c, h.service.StopServer)
}
