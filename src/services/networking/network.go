package networking

import (
	"context"

	"github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
	"github.com/itera-io/openstack-web-client/api/dto"
	"github.com/itera-io/openstack-web-client/pkg/logging"
)

// GetNetworks retrieves the networks.
func (s *Service) GetNetworks(ctx context.Context, req *dto.ListNetworkRequest, authUtils *dto.AuthUtils) (*dto.ListNetworkResponse, error) {
	client, err := s.newNetworkV2Client(authUtils)
	if err != nil {
		return nil, err
	}
	listOpts := networks.ListOpts{}
	allPages, err := networks.List(client, listOpts).AllPages()
	if err != nil {
		s.Logger.Error(logging.IdentityClient, logging.ExternalService, "Failed to get networks", nil)
		return nil, err
	}
	networks, err := networks.ExtractNetworks(allPages)
	if err != nil {
		return nil, err
	}
	return &dto.ListNetworkResponse{Networks: networks}, nil
}
