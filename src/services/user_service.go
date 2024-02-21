package services

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/itera-io/openstack-web-client/api/dto"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

// Validate user
func (s *UserService) ValidateUser(req *dto.ValidateUserRequest) (bool, error) {
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
