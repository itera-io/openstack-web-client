package dto

import "github.com/gophercloud/gophercloud/openstack/sharedfilesystems/v2/availabilityzones"

type ListAvailabilityZoneRequest struct{}

type ListAvailabilityZoneResponse struct {
	AvailabilityZones []availabilityzones.AvailabilityZone `json:"availabilityzones"`
}
