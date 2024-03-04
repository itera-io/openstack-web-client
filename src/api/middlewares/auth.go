package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itera-io/openstack-web-client/api/helper"
	"github.com/itera-io/openstack-web-client/config"
	"github.com/itera-io/openstack-web-client/constants"
	"github.com/itera-io/openstack-web-client/pkg/service_errors"
)

func Authentication(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		token := c.GetHeader(constants.TokenKey)
		if token == "" {
			err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenRequired}
		}
		url := c.GetHeader(constants.AuthUrlKey)
		if url == "" {
			err = &service_errors.ServiceError{EndUserMessage: service_errors.BaseUrlRequired}
		}
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.GenerateBaseResponseWithError(
				nil, false, helper.AuthError, err,
			))
			return
		}
		c.Set(constants.TokenKey, token)
		c.Set(constants.AuthUrlKey, url)
		c.Next()
	}
}
