package compute

import (
	"context"

	"github.com/gophercloud/gophercloud/openstack/compute/v2/flavors"
	"github.com/itera-io/openstack-web-client/api/dto"
	"github.com/itera-io/openstack-web-client/pkg/logging"
)

func (s *Service) ListFlavors(ctx context.Context, req *dto.ListFlavorRequest, authUtils *dto.AuthUtils) (*dto.ListFlavorResponse, error) {
	client, err := s.newNewComputeV2(authUtils)
	if err != nil {
		return nil, err
	}
	listOpts := flavors.ListOpts{
		AccessType: flavors.PublicAccess,
		MinDisk:    req.MinDisk,
		MinRAM:     req.MinRAM,
	}

	allPages, err := flavors.ListDetail(client, listOpts).AllPages()
	if err != nil {
		s.Logger.Error(logging.IdentityClient, logging.ExternalService, "Failed to list flavors", nil)
		return nil, err
	}

	allFlavors, err := flavors.ExtractFlavors(allPages)
	if err != nil {
		return nil, err
	}

	return &dto.ListFlavorResponse{Flavors: allFlavors}, nil
}

func (s *Service) GetFlavor(ctx context.Context, flavorId string, authUtils *dto.AuthUtils) (*dto.GetFlavorResponse, error) {
	client, err := s.newNewComputeV2(authUtils)
	if err != nil {
		return nil, err
	}
	getResult := flavors.Get(client, flavorId)

	flavor, err := getResult.Extract()
	if err != nil {
		s.Logger.Error(logging.IdentityClient, logging.ExternalService, "Failed to get flavor", nil)
		return nil, err
	}
	return &dto.GetFlavorResponse{Flavor: *flavor}, nil
}
