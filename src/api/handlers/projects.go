package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/itera-io/openstack-web-client/api/dto"
	_ "github.com/itera-io/openstack-web-client/api/helper"
	"github.com/itera-io/openstack-web-client/config"
	"github.com/itera-io/openstack-web-client/services"
)

type ProjectsHandler struct {
	service *services.ProjectService
}

func NewProjectsHandler(cfg *config.Config) *ProjectsHandler {
	service := services.NewProjectService(cfg)
	return &ProjectsHandler{service: service}
}

// ListProject godoc
// @Summary List Project
// @Description List Project
// @Tags Projects
// @Accept  json
// @Produce  json
// @Param Request body dto.ListProjectRequest true "ListProjectRequest"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.ListProjectResponse} "ListProject response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 401 {object} helper.BaseHttpResponse "Unauthorized request"
// @Router /v3/Projects [get]
// @Security AuthBearer
func (h *ProjectsHandler) ListProjects(c *gin.Context) {
	GetByFilter(c, h.service.ListProjects)
}
