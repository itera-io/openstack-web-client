package dto

import "github.com/gophercloud/gophercloud/openstack/blockstorage/extensions/limits"

type GetVolumeLimitRequest struct{}

type GetVolumeLimitResponse struct {
	VolumeLimits limits.Limits `json:"limits"`
}
