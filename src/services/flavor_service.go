package services

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/flavors"
	"github.com/itera-io/openstack-web-client/api/dto"
	"github.com/itera-io/openstack-web-client/config"
	"github.com/itera-io/openstack-web-client/pkg/logging"
)

type FlavorService struct{ Logger logging.Logger }

func NewFlavorService(cfg *config.Config) *FlavorService {
	return &FlavorService{Logger: logging.NewLogger(cfg)}
}

func (s *FlavorService) ListFlavors(req *dto.ListFlavorRequest, authUtils *dto.AuthUtils) (*dto.ListFlavorResponse, error) {
	r := &dto.ListFlavorResponse{}
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: authUtils.BaseUrl,
		TokenID:          authUtils.Token,
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		return nil, err
	}
	computeClient, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		s.Logger.Error(logging.IdentityClient, logging.ExternalService, err.Error(), nil)
		return nil, err
	}
	listOpts := flavors.ListOpts{
		AccessType: flavors.PublicAccess,
	}

	allPages, err := flavors.ListDetail(computeClient, listOpts).AllPages()
	if err != nil {
		return nil, err
	}

	allFlavors, err := flavors.ExtractFlavors(allPages)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(allFlavors); i++ {
		r.Flavors = append(r.Flavors, dto.Flavor{Id: allFlavors[i].ID, Name: allFlavors[i].Name})
	}
	return r, nil
}
