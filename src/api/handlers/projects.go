package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itera-io/openstack-web-client/api/dto"
	"github.com/itera-io/openstack-web-client/api/helper"
	"github.com/itera-io/openstack-web-client/config"
	"github.com/itera-io/openstack-web-client/constants"
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
	// req := new(dto.ListProjectRequest)
	// err := c.ShouldBindJSON(&req)
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest,
	// 		helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
	// 	return
	// }
	var t, _ = c.Get(constants.TokenKey)
	var u, _ = c.Get(constants.AuthUrlKey)
	authUtils := &dto.AuthUtils{Token: t.(string), BaseUrl: u.(string)}
	res, err := h.service.ListProjects(new(dto.ListProjectRequest), authUtils)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, helper.Success))
}
