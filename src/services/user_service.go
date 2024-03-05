package services

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/projects"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/tokens"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/users"
	"github.com/itera-io/openstack-web-client/api/dto"
	"github.com/itera-io/openstack-web-client/config"
	"github.com/itera-io/openstack-web-client/pkg/logging"
)

type UserService struct{ Logger logging.Logger }

func NewUserService(cfg *config.Config) *UserService {
	return &UserService{Logger: logging.NewLogger(cfg)}
}

// Validate user
func (s *UserService) Validate(req *dto.ValidateUserRequest) (bool, error) {
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: req.Url,
		Username:         req.Username,
		Password:         req.Password,
		DomainName:       req.Domain,
		AllowReauth:      true,
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		return false, err
	}

	// Example of creating a service client, e.g., for Compute
	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		return false, err
	}

	// Perform a test operation to validate if authentication is successful
	_, err = servers.List(client, servers.ListOpts{}).AllPages()
	if err != nil {
		return false, err
	}

	return true, nil
}

// Authenticate user
func (s *UserService) Authenticate(req *dto.AuthenticateUserRequest) (*dto.AuthenticateUserResponse, error) {
	r := &dto.AuthenticateUserResponse{}
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: req.Url,
		Username:         req.Username,
		Password:         req.Password,
		DomainName:       req.Domain,
		AllowReauth:      true,
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		return nil, err
	}

	identityClient, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})
	if err != nil {
		return nil, err
	}
	result := tokens.Create(identityClient, &opts)
	if result.Err != nil {
		return nil, result.Err
	}
	result.ExtractInto(&r.Token)
	tokenId, _ := result.ExtractTokenID()
	r.TokenID = tokenId
	return r, nil
}

// List users' projects
func (s *UserService) ListUserProjects(userId string, authUtils *dto.AuthUtils) (*dto.ListUserProjectResponse, error) {
	r := &dto.ListUserProjectResponse{}
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

	allPages, err := users.ListProjects(identityClient, userId).AllPages()
	if err != nil {
		return nil, err
	}

	allProjects, err := projects.ExtractProjects(allPages)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(allProjects); i++ {
		r.UserProjects = append(r.UserProjects, dto.CommonDto{Id: allProjects[i].ID, Name: allProjects[i].Name})
	}
	return r, nil
}
