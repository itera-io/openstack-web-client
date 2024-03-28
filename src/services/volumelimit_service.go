package services

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/extensions/limits"
	"github.com/itera-io/openstack-web-client/api/dto"
	"github.com/itera-io/openstack-web-client/config"
	"github.com/itera-io/openstack-web-client/pkg/logging"
)

type VolumeLimitService struct{ Logger logging.Logger }

func NewVolumeLimitService(cfg *config.Config) *VolumeLimitService {
	return &VolumeLimitService{Logger: logging.NewLogger(cfg)}
}

func (s *VolumeLimitService) GetLimits(ctx context.Context, req *dto.GetVolumeLimitRequest, authUtils *dto.AuthUtils) (*dto.GetVolumeLimitResponse, error) {
	r := &dto.GetVolumeLimitResponse{}
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: authUtils.BaseUrl,
		TokenID:          authUtils.Token,
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		return nil, err
	}
	client, err := openstack.NewBlockStorageV3(provider, gophercloud.EndpointOpts{})
	if err != nil {
		s.Logger.Error(logging.IdentityClient, logging.ExternalService, err.Error(), nil)
		return nil, err
	}
	limits, err := limits.Get(client).Extract()
	if err != nil {
		return nil, err
	}

	r.VolumeLimits = *limits
	return r, nil
}
