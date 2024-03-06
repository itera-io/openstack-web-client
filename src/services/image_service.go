package services

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/images"
	"github.com/itera-io/openstack-web-client/api/dto"
	"github.com/itera-io/openstack-web-client/config"
	"github.com/itera-io/openstack-web-client/pkg/logging"
)

type ImageService struct{ Logger logging.Logger }

func NewImageService(cfg *config.Config) *ImageService {
	return &ImageService{Logger: logging.NewLogger(cfg)}
}

func (s *ImageService) ListImages(ctx context.Context, req *dto.ListImageRequest, authUtils *dto.AuthUtils) (*dto.ListImageResponse, error) {
	r := &dto.ListImageResponse{}
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: authUtils.BaseUrl,
		TokenID:          authUtils.Token,
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		return nil, err
	}
	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		s.Logger.Error(logging.IdentityClient, logging.ExternalService, err.Error(), nil)
		return nil, err
	}
	listOpts := images.ListOpts{}

	allPages, err := images.ListDetail(client, listOpts).AllPages()
	if err != nil {
		return nil, err
	}

	allImages, err := images.ExtractImages(allPages)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(allImages); i++ {
		r.Images = append(r.Images, dto.CommonDto{Id: allImages[i].ID, Name: allImages[i].Name})
	}
	return r, nil
}

func (s *ImageService) GetImage(ctx context.Context, id string, authUtils *dto.AuthUtils) (*dto.GetImageResponse, error) {
	r := &dto.GetImageResponse{}
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: authUtils.BaseUrl,
		TokenID:          authUtils.Token,
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		return nil, err
	}
	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		s.Logger.Error(logging.IdentityClient, logging.ExternalService, err.Error(), nil)
		return nil, err
	}
	getResult := images.Get(client, id)

	image, err := getResult.Extract()
	if err != nil {
		return nil, err
	}

	r.Image = append(r.Image, dto.ImageDto{ID: image.ID, Name: image.Name, MinDisk: image.MinDisk, MinRAM: image.MinRAM, Status: image.Status})

	return r, nil
}
