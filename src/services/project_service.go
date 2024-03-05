package services

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/projects"
	"github.com/itera-io/openstack-web-client/api/dto"
	"github.com/itera-io/openstack-web-client/config"
	"github.com/itera-io/openstack-web-client/pkg/logging"
)

type ProjectService struct{ Logger logging.Logger }

func NewProjectService(cfg *config.Config) *ProjectService {
	return &ProjectService{Logger: logging.NewLogger(cfg)}
}

func (s *ProjectService) ListProjects(ctx context.Context, req *dto.ListProjectRequest, authUtils *dto.AuthUtils) (*dto.ListProjectResponse, error) {
	r := &dto.ListProjectResponse{}
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: authUtils.BaseUrl,
		TokenID:          authUtils.Token,
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		return nil, err
	}
	identityClient, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})
	if err != nil {
		s.Logger.Error(logging.IdentityClient, logging.ExternalService, err.Error(), nil)
		return nil, err
	}
	listOpts := projects.ListOpts{
		Enabled: gophercloud.Enabled,
	}

	allPages, err := projects.List(identityClient, listOpts).AllPages()
	if err != nil {
		return nil, err
	}

	allProjects, err := projects.ExtractProjects(allPages)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(allProjects); i++ {
		r.Projects = append(r.Projects, dto.CommonDto{Id: allProjects[i].ID, Name: allProjects[i].Name})
	}
	return r, nil
}
