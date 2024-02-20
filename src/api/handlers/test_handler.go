package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itera-io/openstack-web-client/api/helper"
)

type header struct {
	UserId  string
	Browser string
}

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) Test(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"result": "Test",
	})
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("Test", true, 0))

}

func (h *TestHandler) Users(c *gin.Context) {

	c.JSON(http.StatusOK, helper.GenerateBaseResponse("Users", true, 0))

}
