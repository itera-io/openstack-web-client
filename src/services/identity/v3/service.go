package v3

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

func (s *Service) newIdetityV3Client(authUtils *dto.AuthUtils) (*gophercloud.ServiceClient, error) {
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: authUtils.BaseUrl,
		TokenID:          authUtils.Token,
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		s.Logger.Error(logging.IdentityClient, logging.ExternalService, "Failed to authenticate client", nil)
		return nil, err
	}
	client, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})
	if err != nil {
		s.Logger.Error(logging.IdentityClient, logging.ExternalService, "Failed to create identity V3 client", nil)
		return nil, err
	}

	return client, nil
}
