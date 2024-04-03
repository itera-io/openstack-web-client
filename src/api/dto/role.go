package dto

import "github.com/gophercloud/gophercloud/openstack/identity/v3/roles"

type ListRoleRequest struct {
	// DomainID filters the response by a domain ID.
	DomainID string `form:"domain_id"`
	// Name filters the response by role name.
	Name string `form:"name"`
}

type ListRoleResponse struct {
	Roles []roles.Role `json:"roles"`
}

type AssignRoleRequest struct {
	UserId    string `json:"userId"`
	ProjectId string `json:"projectId"`
}

type AssignRoleResponse struct{}
