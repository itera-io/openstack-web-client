package dto

import (
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)

type ListServerRequest struct {

	// Flavor is the name of the flavor in URL format.
	Flavor string `form:"flavor"`

	// IP is a regular expression to match the IPv4 address of the server.
	IP string `form:"ip"`

	// Name of the server as a string; can be queried with regular expressions.
	// Realize that ?name=bob returns both bob and bobb. If you need to match bob
	// only, you can use a regular expression matching the syntax of the
	// underlying database server implemented for Compute.
	Name string `form:"name"`

	// Status is the value of the status of the server so that you can filter on
	// "ACTIVE" for example.
	Status string `form:"status"`

	// Limit is an integer value for the limit of values to return.
	Limit int `form:"limit"`

	// TenantID lists servers for a particular tenant.
	// Setting "AllTenants = true" is required.
	TenantID string `form:"tenant_id"`

	// This requires the client to be set to microversion 2.83 or later, unless
	// the user is an admin.
	// UserID lists servers for a particular user.
	UserID string `form:"user_id"`

	// This requires the client to be set to microversion 2.26 or later.
	// Tags filters on specific server tags. All tags must be present for the server.
	Tags string `form:"tags"`
}

type ListServerResponse struct {
	Servers []servers.Server `json:"servers"`
}
type GetServerResponse struct {
	Server servers.Server `json:"server"`
}
