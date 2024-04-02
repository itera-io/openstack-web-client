package v3

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/projects"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/regions"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/tokens"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/users"
	"github.com/itera-io/openstack-web-client/api/dto"
)

// Validate user
func (s *Service) Validate(ctx context.Context, req *dto.ValidateUserRequest) (*dto.ValidateUserResponse, error) {
	r := &dto.ValidateUserResponse{IsAuthenticated: false}
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: req.Url,
		Username:         req.Username,
		Password:         req.Password,
		DomainName:       req.Domain,
		AllowReauth:      true,
	}

	client, err := s.newIdetityV3ClientByAuthOpts(opts)
	if err != nil {
		return nil, err
	}

	// Perform a test operation to validate if authentication is successful
	_, err = regions.List(client, regions.ListOpts{}).AllPages()
	if err != nil {
		return nil, err
	}
	r.IsAuthenticated = true
	return r, nil
}

// Authenticate user
func (s *Service) Authenticate(ctx context.Context, req *dto.AuthenticateUserRequest) (*dto.AuthenticateUserResponse, error) {
	r := &dto.AuthenticateUserResponse{}
	opts := gophercloud.AuthOptions{
		IdentityEndpoint:            req.Url,
		Username:                    req.Username,
		Password:                    req.Password,
		DomainName:                  req.Domain,
		AllowReauth:                 true,
		ApplicationCredentialID:     req.ApplicationCredentialID,
		ApplicationCredentialName:   req.ApplicationCredentialName,
		ApplicationCredentialSecret: req.ApplicationCredentialSecret,
	}
	client, err := s.newIdetityV3ClientByAuthOpts(opts)

	if err != nil {
		return nil, err
	}
	result := tokens.Create(client, &opts)
	if result.Err != nil {
		return nil, result.Err
	}
	result.ExtractInto(&r.Token)
	tokenId, _ := result.ExtractTokenID()
	r.TokenID = tokenId
	return r, nil
}

// List users' projects
func (s *Service) ListUserProjects(ctx context.Context, userId string, authUtils *dto.AuthUtils) (*dto.ListUserProjectResponse, error) {
	r := &dto.ListUserProjectResponse{}
	client, err := s.newIdetityV3Client(authUtils)

	if err != nil {
		return nil, err
	}

	allPages, err := users.ListProjects(client, userId).AllPages()
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

// Create a user.
func (s Service) CreateUser(ctx context.Context, req *dto.CreateUserRequest, authUtils *dto.AuthUtils) (*dto.CreateUserResponse, error) {
	client, err := s.newIdetityV3Client(authUtils)
	if err != nil {
		return nil, err
	}
	opts := users.CreateOpts{
		DomainID:    req.DomainID,
		Enabled:     req.Enabled,
		Name:        req.Name,
		Password:    req.Password,
		Description: req.Description,
	}
	result := users.Create(client, opts)
	if result.Err != nil {
		return nil, result.Err
	}
	res, _ := result.Extract()
	return &dto.CreateUserResponse{User: *res}, nil
}

// List all users
func (s *Service) ListUsers(ctx context.Context, req *dto.ListUserRequest, authUtils *dto.AuthUtils) (*dto.ListUserResponse, error) {
	client, err := s.newIdetityV3Client(authUtils)
	if err != nil {
		return nil, err
	}
	listOpts := users.ListOpts{
		Enabled:  req.Enabled,
		DomainID: req.DomainID,
		Name:     req.Name,
	}

	allPages, err := users.List(client, listOpts).AllPages()
	if err != nil {
		return nil, err
	}

	allUsers, err := users.ExtractUsers(allPages)
	if err != nil {
		return nil, err
	}
	return &dto.ListUserResponse{Users: allUsers}, nil
}
