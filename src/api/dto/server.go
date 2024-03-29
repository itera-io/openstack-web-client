package dto

import (
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)

type ListServerRequest struct {
	MinDisk int `form:"minDisk"`
	MinRAM  int `form:"minRam"`
}

type ListServerResponse struct {
	Servers []servers.Server `json:"servers"`
}
type GetServerResponse struct {
	Server servers.Server `json:"server"`
}
