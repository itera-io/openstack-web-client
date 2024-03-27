package dto

import "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/quotas"

type GetNetworkQuotaRequest struct {
	ProjectID string `form:"projectId"`
}

type GetNetworkQuotaResponse struct {
	NetworkQuota quotas.Quota `json:"quota"`
}
