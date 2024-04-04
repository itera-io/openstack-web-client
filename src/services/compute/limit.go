package compute

import (
	"context"

	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/limits"
	"github.com/itera-io/openstack-web-client/api/dto"
	"github.com/itera-io/openstack-web-client/pkg/logging"
)

func (s *Service) GetLimits(ctx context.Context, req *dto.GetComputeLimitRequest, authUtils *dto.AuthUtils) (*dto.GetComputeLimitResponse, error) {
	client, err := s.newNewComputeV2(authUtils)
	if err != nil {
		return nil, err
	}
	getOpts := limits.GetOpts{
		TenantID: req.TenantID,
	}
	limits, err := limits.Get(client, getOpts).Extract()
	if err != nil {
		s.Logger.Error(logging.ComputeClient, logging.ExternalService, "Failed to get compute limits", nil)
		return nil, err
	}
	return &dto.GetComputeLimitResponse{ComputeLimits: *limits}, nil
}
