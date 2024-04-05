package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/itera-io/openstack-web-client/api/dto"
	_ "github.com/itera-io/openstack-web-client/api/helper"
	"github.com/itera-io/openstack-web-client/config"
	"github.com/itera-io/openstack-web-client/services/imageservice"
)

type ImagesHandler struct {
	service *imageservice.Service
}

func NewImagesHandler(cfg *config.Config) *ImagesHandler {
	service := imageservice.NewService(cfg)
	return &ImagesHandler{service: service}
}

// ListImage godoc
// @Summary List Image
// @Description List Image
// @Tags Images
// @Accept  json
// @Produce  json
// @Param Request body dto.ListImageRequest true "ListImage Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.ListImageResponse} "ListImage response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 401 {object} helper.BaseHttpResponse "Unauthorized request"
// @Router /v2/images [get]
// @Security AuthBearer
func (h *ImagesHandler) ListImages(c *gin.Context) {
	GetByFilter(c, h.service.ListImages)
}

// GetImage godoc
// @Summary Get Image
// @Description Get Image by ID
// @Tags Images
// @Accept  json
// @Produce  json
// @Param id path string true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.GetImageResponse} "GetImage response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 401 {object} helper.BaseHttpResponse "Unauthorized request"
// @Router /v2/images/{id} [get]
// @Security AuthBearer
func (h *ImagesHandler) GetImage(c *gin.Context) {
	GetById(c, h.service.GetImage)
}
