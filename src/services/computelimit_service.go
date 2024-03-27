package services

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/limits"
	"github.com/itera-io/openstack-web-client/api/dto"
	"github.com/itera-io/openstack-web-client/config"
	"github.com/itera-io/openstack-web-client/pkg/logging"
)

type ComputeLimitService struct{ Logger logging.Logger }

func NewComputeLimitService(cfg *config.Config) *ComputeLimitService {
	return &ComputeLimitService{Logger: logging.NewLogger(cfg)}
}

func (s *ComputeLimitService) Get(ctx context.Context, req *dto.GetComputeLimitRequest, authUtils *dto.AuthUtils) (*dto.GetComputeLimitResponse, error) {
	r := &dto.GetComputeLimitResponse{}
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: authUtils.BaseUrl,
		TokenID:          authUtils.Token,
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		return nil, err
	}
	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		s.Logger.Error(logging.IdentityClient, logging.ExternalService, err.Error(), nil)
		return nil, err
	}
	getOpts := limits.GetOpts{
		TenantID: req.TenantID,
	}
	limits, err := limits.Get(client, getOpts).Extract()
	if err != nil {
		return nil, err
	}

	r.ComputeLimits = *limits
	return r, nil
}
