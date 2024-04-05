package imageservice

import (
	"context"

	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
	"github.com/itera-io/openstack-web-client/api/dto"
)

func (s *Service) ListImages(ctx context.Context, req *dto.ListImageRequest, authUtils *dto.AuthUtils) (*dto.ListImageResponse, error) {
	client, err := s.newImageServiceV2(authUtils)
	if err != nil {
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
	r := &dto.ListImageResponse{}
	for _, image := range allImages {
		r.Images = append(r.Images, dto.ImageListItem{ID: image.ID, Name: image.Name, Tags: image.Tags})
	}
	return r, nil
}

func (s *Service) GetImage(ctx context.Context, id string, authUtils *dto.AuthUtils) (*dto.GetImageResponse, error) {
	client, err := s.newImageServiceV2(authUtils)
	if err != nil {
		return nil, err
	}
	getResult := images.Get(client, id)

	image, err := getResult.Extract()
	if err != nil {
		return nil, err
	}
	r := &dto.GetImageResponse{}
	r.Image = append(r.Image, dto.ImageDto{ID: image.ID, Name: image.Name, MinDiskGigabytes: image.MinDiskGigabytes, MinRAMMegabytes: image.MinRAMMegabytes, Status: string(image.Status)})

	return r, nil
}
