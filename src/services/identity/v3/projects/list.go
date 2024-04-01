package projects

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/projects"
	"github.com/itera-io/openstack-web-client/api/dto"
)

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
