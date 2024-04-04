package v3

import (
	"context"

	"github.com/gophercloud/gophercloud/openstack/identity/v3/roles"
	"github.com/itera-io/openstack-web-client/api/dto"
)

// List all roles
func (s *Service) ListRoles(ctx context.Context, req *dto.ListRoleRequest, authUtils *dto.AuthUtils) (*dto.ListRoleResponse, error) {
	client, err := s.newIdetityV3Client(authUtils)
	if err != nil {
		return nil, err
	}
	listOpts := roles.ListOpts{
		DomainID: req.DomainID,
		Name:     req.Name,
	}

	allPages, err := roles.List(client, listOpts).AllPages()
	if err != nil {
		return nil, err
	}

	allRoles, err := roles.ExtractRoles(allPages)
	if err != nil {
		return nil, err
	}
	return &dto.ListRoleResponse{Roles: allRoles}, nil
}

// Assign a Role to a User in a Project.
func (s *Service) AssignRole(ctx context.Context, roleId string, req *dto.AssignRoleRequest, authUtils *dto.AuthUtils) (*dto.AssignRoleResponse, error) {
	client, err := s.newIdetityV3Client(authUtils)
	if err != nil {
		return nil, err
	}
	opts := roles.AssignOpts{
		UserID:    req.UserId,
		ProjectID: req.ProjectId,
	}
	err = roles.Assign(client, roleId, opts).ExtractErr()
	if err != nil {
		return nil, err
	}
	return &dto.AssignRoleResponse{}, nil
}
