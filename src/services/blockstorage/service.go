package blockstorage

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

// newBlockStorageClient creates a new authenticated block storage client.
func (s *Service) newBlockStorageClient(authUtils *dto.AuthUtils) (*gophercloud.ServiceClient, error) {
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: authUtils.BaseUrl,
		TokenID:          authUtils.Token,
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		s.Logger.Error(logging.BlockStorageClient, logging.ExternalService, "Failed to authenticate client", nil)
		return nil, err
	}

	client, err := openstack.NewBlockStorageV3(provider, gophercloud.EndpointOpts{})
	if err != nil {
		s.Logger.Error(logging.BlockStorageClient, logging.ExternalService, "Failed to create block storage V3 client", nil)
		return nil, err
	}

	return client, nil
}
