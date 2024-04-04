package compute

import (
	"context"

	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/shelveunshelve"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/startstop"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/itera-io/openstack-web-client/api/dto"
	"github.com/itera-io/openstack-web-client/pkg/logging"
)

// Get server by id.
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

// List all servers.
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

// Reboot server by id.
func (s *Service) RebootServer(ctx context.Context, serverId string, req *dto.RebootServerRequest, authUtils *dto.AuthUtils) (*dto.RebootServerResponse, error) {
	client, err := s.newNewComputeV2(authUtils)
	if err != nil {
		return nil, err
	}
	opts := servers.RebootOpts{Type: servers.RebootMethod(req.Type)}
	err = servers.Reboot(client, serverId, opts).ExtractErr()
	if err != nil {
		s.Logger.Error(logging.ComputeClient, logging.ExternalService, "Failed to reboot server", nil)
		return nil, err
	}
	return &dto.RebootServerResponse{}, nil
}

// Start server by id.
func (s *Service) StartServer(ctx context.Context, serverId string, authUtils *dto.AuthUtils) error {
	client, err := s.newNewComputeV2(authUtils)
	if err != nil {
		return err
	}
	err = startstop.Start(client, serverId).ExtractErr()
	if err != nil {
		s.Logger.Error(logging.ComputeClient, logging.ExternalService, "Failed to start server", nil)
		return err
	}
	return nil
}

// Stop server by id.
func (s *Service) StopServer(ctx context.Context, serverId string, authUtils *dto.AuthUtils) error {
	client, err := s.newNewComputeV2(authUtils)
	if err != nil {
		return err
	}
	err = startstop.Stop(client, serverId).ExtractErr()
	if err != nil {
		s.Logger.Error(logging.ComputeClient, logging.ExternalService, "Failed to stop server", nil)
		return err
	}
	return nil
}

// Shelve server by id.
func (s *Service) ShelveServer(ctx context.Context, serverId string, authUtils *dto.AuthUtils) error {
	client, err := s.newNewComputeV2(authUtils)
	if err != nil {
		return err
	}
	err = shelveunshelve.Shelve(client, serverId).ExtractErr()
	if err != nil {
		s.Logger.Error(logging.ComputeClient, logging.ExternalService, "Failed to shelve server", nil)
		return err
	}
	return nil
}

// Unshelve server by id.
func (s *Service) UnshelveServer(ctx context.Context, serverId string, authUtils *dto.AuthUtils) error {
	client, err := s.newNewComputeV2(authUtils)
	if err != nil {
		return err
	}
	opts := shelveunshelve.UnshelveOpts{}
	err = shelveunshelve.Unshelve(client, serverId, opts).ExtractErr()
	if err != nil {
		s.Logger.Error(logging.ComputeClient, logging.ExternalService, "Failed to unshelve server", nil)
		return err
	}
	return nil
}
