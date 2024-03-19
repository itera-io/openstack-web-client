package dto

import "github.com/gophercloud/gophercloud/openstack/compute/apiversions"

type ListApiVersionRequest struct{}

type ListApiVersionResponse struct {
	ApiVersions []apiversions.APIVersion `json:"apiversions"`
}
