package networking

import (
	"context"

	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/quotas"
	"github.com/itera-io/openstack-web-client/api/dto"
)

// GetQuotas retrieves the available networking quotas.
func (s *Service) GetQuotas(ctx context.Context, req *dto.GetNetworkQuotaRequest, authUtils *dto.AuthUtils) (*dto.GetNetworkQuotaResponse, error) {
	client, err := s.newNetworkV2Client(authUtils)
	if err != nil {
		return nil, err
	}

	quota, err := quotas.Get(client, req.ProjectID).Extract()
	if err != nil {
		return nil, err
	}

	return &dto.GetNetworkQuotaResponse{NetworkQuota: *quota}, nil
}
