package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itera-io/openstack-web-client/api/dto"
	"github.com/itera-io/openstack-web-client/api/helper"
	"github.com/itera-io/openstack-web-client/services"
)

type UsersHandler struct {
	service *services.UserService
}

func NewUsersHandler() *UsersHandler {
	service := services.NewUserService()
	return &UsersHandler{service: service}
}

// ValidateUser godoc
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
func (h *UsersHandler) ValidateUser(c *gin.Context) {
	req := new(dto.ValidateUserRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}
	authenticated, err := h.service.ValidateUser(req)
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
