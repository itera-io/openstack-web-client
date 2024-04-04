package networking

import (
	"context"

	"github.com/gophercloud/gophercloud/openstack/networking/v2/subnets"
	"github.com/itera-io/openstack-web-client/api/dto"
	"github.com/itera-io/openstack-web-client/pkg/logging"
)

// GetSubnets retrieves the subnets.
func (s *Service) GetSubnets(ctx context.Context, req *dto.ListNetworkSubnetRequest, authUtils *dto.AuthUtils) (*dto.ListNetworkSubnetResponse, error) {
	client, err := s.newNetworkV2Client(authUtils)
	if err != nil {
		return nil, err
	}
	listOpts := subnets.ListOpts{
		ProjectID: req.ProjectID,
	}
	allPages, err := subnets.List(client, listOpts).AllPages()
	if err != nil {
		s.Logger.Error(logging.NetworkClient, logging.ExternalService, "Failed to get subnets", nil)
		return nil, err
	}
	subnetResult, err := subnets.ExtractSubnets(allPages)
	if err != nil {
		return nil, err
	}
	return &dto.ListNetworkSubnetResponse{Subnets: subnetResult}, nil
}
