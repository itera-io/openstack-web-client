package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/itera-io/openstack-web-client/api/dto"
	_ "github.com/itera-io/openstack-web-client/api/helper"
	"github.com/itera-io/openstack-web-client/config"
	"github.com/itera-io/openstack-web-client/services"
)

type UsersHandler struct {
	service *services.UserService
}

func NewUsersHandler(cfg *config.Config) *UsersHandler {
	service := services.NewUserService(cfg)
	return &UsersHandler{service: service}
}

// Validate godoc
// @Summary Validate User
// @Description Validate User
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Request body dto.ValidateUserRequest true "ValidateUserRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/users/validate [post]
func (h *UsersHandler) Validate(c *gin.Context) {
	Create(c, h.service.Validate)
}

// Authenticate godoc
// @Summary Authenticate User
// @Description Authenticate User
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Request body dto.AuthenticateUserRequest true "AuthenticateUserRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/users/auth [post]
func (h *UsersHandler) Authenticate(c *gin.Context) {
	Create(c, h.service.Authenticate)
}

// ListUserProject godoc
// @Summary List UserProject
// @Description List UserProject
// @Tags UserProjects
// @Accept  json
// @Produce  json
// @Param id path string true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.ListUserProjectResponse} "ListUserProject response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 401 {object} helper.BaseHttpResponse "Unauthorized request"
// @Router /v3/users/{id}/projects [get]
// @Security AuthBearer
func (h *UsersHandler) ListUserProjects(c *gin.Context) {
	GetById(c, h.service.ListUserProjects)
}
