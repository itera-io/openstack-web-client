package dto

import "github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumetypes"

type ListVolumeTypeRequest struct{}

type ListVolumeTypeResponse struct {
	VolumeTypes []volumetypes.VolumeType `json:"volumetypes"`
}
