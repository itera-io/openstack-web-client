package services

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/apiversions"
	"github.com/itera-io/openstack-web-client/api/dto"
	"github.com/itera-io/openstack-web-client/config"
	"github.com/itera-io/openstack-web-client/pkg/logging"
)

type ApiVersionService struct{ Logger logging.Logger }

func NewApiVersionService(cfg *config.Config) *ApiVersionService {
	return &ApiVersionService{Logger: logging.NewLogger(cfg)}
}

func (s *ApiVersionService) ListApiVersions(ctx context.Context, req *dto.ListApiVersionRequest, authUtils *dto.AuthUtils) (*dto.ListApiVersionResponse, error) {
	r := &dto.ListApiVersionResponse{}
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: authUtils.BaseUrl,
		TokenID:          authUtils.Token,
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		return nil, err
	}
	computeClient, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		s.Logger.Error(logging.IdentityClient, logging.ExternalService, err.Error(), nil)
		return nil, err
	}

	allPages, err := apiversions.List(computeClient).AllPages()
	if err != nil {
		return nil, err
	}

	allApiVersions, err := apiversions.ExtractAPIVersions(allPages)
	if err != nil {
		return nil, err
	}
	r.ApiVersions = allApiVersions
	return r, nil
}
