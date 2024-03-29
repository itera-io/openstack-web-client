package compute

import (
	"context"

	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/itera-io/openstack-web-client/api/dto"
	"github.com/itera-io/openstack-web-client/pkg/logging"
)

func (s *Service) GetServer(ctx context.Context, serverId string, authUtils *dto.AuthUtils) (*dto.GetServerResponse, error) {
	client, err := s.newNewComputeV2(authUtils)
	if err != nil {
		return nil, err
	}
	getResult := servers.Get(client, serverId)

	server, err := getResult.Extract()
	if err != nil {
		s.Logger.Error(logging.IdentityClient, logging.ExternalService, "Failed to get server", nil)
		return nil, err
	}
	return &dto.GetServerResponse{Server: *server}, nil
}
