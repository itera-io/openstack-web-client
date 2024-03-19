package dto

import "github.com/gophercloud/gophercloud/openstack/compute/apiversions"

type ListApiVersionRequest struct{}

type ListApiVersionResponse struct {
	ApiVersions []apiversions.APIVersion `json:"apiversions"`
}

// type ApiVersionDto struct {
// 	// ID is the unique identifier of the API version.
// 	ID string `json:"id"`

// 	// MinVersion is the minimum microversion supported.
// 	MinVersion string `json:"min_version"`

// 	// Status is the API versions status.
// 	Status string `json:"status"`

// 	// Version is the maximum microversion supported.
// 	Version string `json:"version"`
// }
