package dto

import (
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/quotas"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/subnets"
)

type GetNetworkQuotaRequest struct {
	ProjectID string `form:"projectId"`
}

type GetNetworkQuotaResponse struct {
	NetworkQuota quotas.Quota `json:"quota"`
}

type ListNetworkSubnetRequest struct {
	ProjectID string `form:"projectId"`
}

type ListNetworkSubnetResponse struct {
	Subnets []subnets.Subnet `json:"subnets"`
}

type ListNetworkRequest struct{}

type ListNetworkResponse struct {
	Networks []networks.Network `json:"networks"`
}
