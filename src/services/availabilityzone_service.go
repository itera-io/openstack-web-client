package services

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/sharedfilesystems/v2/availabilityzones"
	"github.com/itera-io/openstack-web-client/api/dto"
	"github.com/itera-io/openstack-web-client/config"
	"github.com/itera-io/openstack-web-client/pkg/logging"
)

type AvailabilityZoneService struct{ Logger logging.Logger }

func NewAvailabilityZoneService(cfg *config.Config) *AvailabilityZoneService {
	return &AvailabilityZoneService{Logger: logging.NewLogger(cfg)}
}

func (s *AvailabilityZoneService) List(ctx context.Context, req *dto.ListAvailabilityZoneRequest, authUtils *dto.AuthUtils) (*dto.ListAvailabilityZoneResponse, error) {
	r := &dto.ListAvailabilityZoneResponse{}
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

	allPages, err := availabilityzones.List(computeClient).AllPages()
	if err != nil {
		return nil, err
	}

	allAvailabilityZones, err := availabilityzones.ExtractAvailabilityZones(allPages)
	if err != nil {
		return nil, err
	}
	r.AvailabilityZones = allAvailabilityZones
	return r, nil
}
