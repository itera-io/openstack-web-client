package compute

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/itera-io/openstack-web-client/api/dto"
	"github.com/itera-io/openstack-web-client/config"
	"github.com/itera-io/openstack-web-client/pkg/logging"
)

type Service struct {
	Logger logging.Logger
}

func NewService(cfg *config.Config) *Service {
	return &Service{Logger: logging.NewLogger(cfg)}
}

// newNewComputeV2 creates a new authenticated compute v2 client.
func (s *Service) newNewComputeV2(authUtils *dto.AuthUtils) (*gophercloud.ServiceClient, error) {
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: authUtils.BaseUrl,
		TokenID:          authUtils.Token,
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		s.Logger.Error(logging.ComputeClient, logging.ExternalService, "Failed to authenticate client", nil)
		return nil, err
	}

	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		s.Logger.Error(logging.ComputeClient, logging.ExternalService, "Failed to create compute V2 client", nil)
		return nil, err
	}

	return client, nil
}
