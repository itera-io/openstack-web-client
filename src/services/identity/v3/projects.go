package v3

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/projects"
	"github.com/itera-io/openstack-web-client/api/dto"
)

func (s Service) CreateProject(ctx context.Context, req *dto.CreateProjectRequest, authUtils *dto.AuthUtils) (*dto.CreateProjectResponse, error) {
	client, err := s.newIdetityV3Client(authUtils)
	if err != nil {
		return nil, err
	}
	opts := projects.CreateOpts{
		DomainID:    req.DomainID,
		Enabled:     req.Enabled,
		IsDomain:    req.IsDomain,
		Name:        req.Name,
		ParentID:    req.ParentID,
		Description: req.Description,
		Tags:        req.Tags,
	}
	result := projects.Create(client, opts)
	if result.Err != nil {
		return nil, result.Err
	}

	res := &dto.CreateProjectResponse{}
	result.ExtractInto(&res.Project)
	return res, nil
}

func (s *Service) ListProjects(ctx context.Context, req *dto.ListProjectRequest, authUtils *dto.AuthUtils) (*dto.ListProjectResponse, error) {
	client, err := s.newIdetityV3Client(authUtils)
	if err != nil {
		return nil, err
	}
	listOpts := projects.ListOpts{
		Enabled: gophercloud.Enabled,
	}

	allPages, err := projects.List(client, listOpts).AllPages()
	if err != nil {
		return nil, err
	}

	allProjects, err := projects.ExtractProjects(allPages)
	if err != nil {
		return nil, err
	}
	return &dto.ListProjectResponse{Projects: allProjects}, nil
}
