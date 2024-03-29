package dto

import "github.com/gophercloud/gophercloud/openstack/compute/v2/flavors"

type ListFlavorRequest struct {
	MinDisk int `form:"minDisk"`
	MinRAM  int `form:"minRam"`
}

type ListFlavorResponse struct {
	Flavors []flavors.Flavor `json:"flavors"`
}
type GetFlavorResponse struct {
	Flavor flavors.Flavor `json:"flavor"`
}
