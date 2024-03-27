package services

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/quotas"
	"github.com/itera-io/openstack-web-client/api/dto"
	"github.com/itera-io/openstack-web-client/config"
	"github.com/itera-io/openstack-web-client/pkg/logging"
)

type NetworkQuotaService struct{ Logger logging.Logger }

func NewNetworkQuotaService(cfg *config.Config) *NetworkQuotaService {
	return &NetworkQuotaService{Logger: logging.NewLogger(cfg)}
}

func (s *NetworkQuotaService) Get(ctx context.Context, req *dto.GetNetworkQuotaRequest, authUtils *dto.AuthUtils) (*dto.GetNetworkQuotaResponse, error) {
	r := &dto.GetNetworkQuotaResponse{}
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: authUtils.BaseUrl,
		TokenID:          authUtils.Token,
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		return nil, err
	}
	client, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		s.Logger.Error(logging.IdentityClient, logging.ExternalService, err.Error(), nil)
		return nil, err
	}
	quota, err := quotas.Get(client, req.ProjectID).Extract()
	if err != nil {
		return nil, err
	}

	r.NetworkQuota = *quota
	return r, nil
}
