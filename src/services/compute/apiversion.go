package compute

import (
	"context"

	"github.com/gophercloud/gophercloud/openstack/compute/apiversions"
	"github.com/itera-io/openstack-web-client/api/dto"
	"github.com/itera-io/openstack-web-client/pkg/logging"
)

func (s *Service) ListApiVersions(ctx context.Context, req *dto.ListApiVersionRequest, authUtils *dto.AuthUtils) (*dto.ListApiVersionResponse, error) {
	client, err := s.newNewComputeV2(authUtils)
	if err != nil {
		return nil, err
	}

	allPages, err := apiversions.List(client).AllPages()
	if err != nil {
		s.Logger.Error(logging.ComputeClient, logging.ExternalService, "Failed to list api versions", nil)
		return nil, err
	}

	allApiVersions, err := apiversions.ExtractAPIVersions(allPages)
	if err != nil {
		return nil, err
	}
	return &dto.ListApiVersionResponse{ApiVersions: allApiVersions}, nil
}
