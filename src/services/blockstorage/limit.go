package blockstorage

import (
	"context"

	"github.com/gophercloud/gophercloud/openstack/blockstorage/extensions/limits"
	"github.com/itera-io/openstack-web-client/api/dto"
	"github.com/itera-io/openstack-web-client/pkg/logging"
)

// GetLimits retrieves the volume limits.
func (s *Service) GetLimits(ctx context.Context, req *dto.GetVolumeLimitRequest, authUtils *dto.AuthUtils) (*dto.GetVolumeLimitResponse, error) {
	client, err := s.newBlockStorageClient(authUtils)
	if err != nil {
		return nil, err
	}

	limitsResult, err := limits.Get(client).Extract()
	if err != nil {
		s.Logger.Error(logging.BlockStorageClient, logging.ExternalService, "Failed to get volume limits", nil)
		return nil, err
	}

	return &dto.GetVolumeLimitResponse{VolumeLimits: *limitsResult}, nil
}
