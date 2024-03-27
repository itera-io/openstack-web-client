package dto

import "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/limits"

type GetComputeLimitRequest struct {
	// The tenant ID to retrieve limits for.
	TenantID string `form:"tenantId"`
}

type GetComputeLimitResponse struct {
	ComputeLimits limits.Limits `json:"limits"`
}
