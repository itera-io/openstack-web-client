package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/itera-io/openstack-web-client/api/dto"
	_ "github.com/itera-io/openstack-web-client/api/helper"
	"github.com/itera-io/openstack-web-client/config"
	v3 "github.com/itera-io/openstack-web-client/services/identity/v3"
)

type ProjectsHandler struct {
	service *v3.Service
}

func NewProjectsHandler(cfg *config.Config) *ProjectsHandler {
	service := v3.NewService(cfg)
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
// @Router /v3/projects [get]
// @Security AuthBearer
func (h *ProjectsHandler) ListProjects(c *gin.Context) {
	GetByFilter(c, h.service.ListProjects)
}

// CreateProject godoc
// @Summary Create Project
// @Description Create Project
// @Tags Projects
// @Accept  json
// @Produce  json
// @Param Request body dto.CreateProjectRequest true "CreateProjectRequest"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CreateProjectResponse} "CreateProject response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 401 {object} helper.BaseHttpResponse "Unauthorized request"
// @Router /v3/projects [post]
// @Security AuthBearer
func (h *ProjectsHandler) CreateProject(c *gin.Context) {
	CreateByAuth(c, h.service.CreateProject)
}
