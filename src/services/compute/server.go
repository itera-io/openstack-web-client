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
		s.Logger.Error(logging.ComputeClient, logging.ExternalService, "Failed to get server", nil)
		return nil, err
	}
	return &dto.GetServerResponse{Server: *server}, nil
}

func (s *Service) ListServers(ctx context.Context, req *dto.ListServerRequest, authUtils *dto.AuthUtils) (*dto.ListServerResponse, error) {
	client, err := s.newNewComputeV2(authUtils)
	if err != nil {
		return nil, err
	}
	listOpts := servers.ListOpts{
		Name:     req.Name,
		Status:   req.Status,
		TenantID: req.TenantID,
		Flavor:   req.Flavor,
		Limit:    req.Limit,
		Tags:     req.Tags,
		UserID:   req.UserID,
		IP:       req.IP,
	}

	allPages, err := servers.List(client, listOpts).AllPages()
	if err != nil {
		s.Logger.Error(logging.ComputeClient, logging.ExternalService, "Failed to list servers", nil)
		return nil, err
	}

	allServers, err := servers.ExtractServers(allPages)
	if err != nil {
		return nil, err
	}

	return &dto.ListServerResponse{Servers: allServers}, nil
}
