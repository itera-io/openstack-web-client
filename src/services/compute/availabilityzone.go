package compute

import (
	"context"

	"github.com/gophercloud/gophercloud/openstack/sharedfilesystems/v2/availabilityzones"
	"github.com/itera-io/openstack-web-client/api/dto"
	"github.com/itera-io/openstack-web-client/pkg/logging"
)

func (s *Service) List(ctx context.Context, req *dto.ListAvailabilityZoneRequest, authUtils *dto.AuthUtils) (*dto.ListAvailabilityZoneResponse, error) {
	client, err := s.newNewComputeV2(authUtils)
	if err != nil {
		return nil, err
	}

	allPages, err := availabilityzones.List(client).AllPages()
	if err != nil {
		s.Logger.Error(logging.ComputeClient, logging.ExternalService, "Failed to list availability zones", nil)
		return nil, err
	}

	allAvailabilityZones, err := availabilityzones.ExtractAvailabilityZones(allPages)
	if err != nil {
		return nil, err
	}
	return &dto.ListAvailabilityZoneResponse{AvailabilityZones: allAvailabilityZones}, nil
}
