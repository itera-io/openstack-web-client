package dto

import (
	"github.com/gophercloud/gophercloud/openstack/blockstorage/extensions/limits"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumetypes"
)

type GetVolumeLimitRequest struct{}

type GetVolumeLimitResponse struct {
	VolumeLimits limits.Limits `json:"limits"`
}

type ListVolumeTypeRequest struct{}

type ListVolumeTypeResponse struct {
	VolumeTypes []volumetypes.VolumeType `json:"volumetypes"`
}
