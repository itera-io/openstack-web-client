package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itera-io/openstack-web-client/api/dto"
	"github.com/itera-io/openstack-web-client/api/helper"
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
	req := new(dto.ValidateUserRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}
	authenticated, err := h.service.Validate(req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	if !authenticated {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			helper.GenerateBaseResponseWithAnyError(nil, false, helper.AuthError, err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(authenticated, true, helper.Success))
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
	req := new(dto.AuthenticateUserRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}
	tDto, err := h.service.Authenticate(req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(tDto, true, helper.Success))
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
