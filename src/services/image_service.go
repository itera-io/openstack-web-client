package services

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
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
	client, err := openstack.NewImageServiceV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		s.Logger.Error(logging.IdentityClient, logging.ExternalService, err.Error(), nil)
		return nil, err
	}
	listOpts := images.ListOpts{
		Owner:  req.Owner,
		Tags:   req.Tags,
		Status: images.ImageStatus(req.Status),
	}

	allPages, err := images.List(client, listOpts).AllPages()
	if err != nil {
		return nil, err
	}

	allImages, err := images.ExtractImages(allPages)
	if err != nil {
		return nil, err
	}

	for _, image := range allImages {
		r.Images = append(r.Images, dto.ImageListItem{ID: image.ID, Name: image.Name, Tags: image.Tags})
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
	client, err := openstack.NewImageServiceV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		s.Logger.Error(logging.IdentityClient, logging.ExternalService, err.Error(), nil)
		return nil, err
	}
	getResult := images.Get(client, id)

	image, err := getResult.Extract()
	if err != nil {
		return nil, err
	}

	r.Image = append(r.Image, dto.ImageDto{ID: image.ID, Name: image.Name, MinDiskGigabytes: image.MinDiskGigabytes, MinRAMMegabytes: image.MinRAMMegabytes, Status: string(image.Status)})

	return r, nil
}
