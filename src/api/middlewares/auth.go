package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/itera-io/openstack-web-client/api/helper"
	"github.com/itera-io/openstack-web-client/config"
	"github.com/itera-io/openstack-web-client/constants"
	"github.com/itera-io/openstack-web-client/pkg/service_errors"
)

func Authentication(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		auth := c.GetHeader(constants.AuthorizationHeaderKey)
		token := strings.Split(auth, " ")
		if auth == "" || token[1] == "" {
			err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenRequired}
		}
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.GenerateBaseResponseWithError(
				nil, false, helper.AuthError, err,
			))
			return
		}
		c.Next()
	}
}
