package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/itera-io/openstack-web-client/api/dto"
	_ "github.com/itera-io/openstack-web-client/api/helper"
	"github.com/itera-io/openstack-web-client/config"
	v3 "github.com/itera-io/openstack-web-client/services/identity/v3"
)

type RolesHandler struct {
	service *v3.Service
}

func NewRolesHandler(cfg *config.Config) *RolesHandler {
	service := v3.NewService(cfg)
	return &RolesHandler{service: service}
}

// ListRole godoc
// @Summary List Role
// @Description List Role
// @Tags Roles
// @Accept  json
// @Produce  json
// @Param Request body dto.ListRoleRequest true "ListRoleRequest"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.ListRoleResponse} "ListRole response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 401 {object} helper.BaseHttpResponse "Unauthorized request"
// @Router /v3/roles [get]
// @Security AuthBearer
func (h *RolesHandler) ListRoles(c *gin.Context) {
	GetByFilter(c, h.service.ListRoles)
}

// AssignRole godoc
// @Summary Assign Role
// @Description Assign Role
// @Tags Roles
// @Accept  json
// @Produce  json
// @Param id path string true "Id"
// @Param Request body dto.AssignRoleRequest true "AssignRoleRequest"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.AssignRoleResponse} "AssignRole response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 401 {object} helper.BaseHttpResponse "Unauthorized request"
// @Router /v3/roles/{id} [put]
// @Security AuthBearer
func (h *RolesHandler) AssignRole(c *gin.Context) {
	Update(c, h.service.AssignRole)
}
