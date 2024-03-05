package services

import (
	"context"

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

func (s *FlavorService) ListFlavors(ctx context.Context, req *dto.ListFlavorRequest, authUtils *dto.AuthUtils) (*dto.ListFlavorResponse, error) {
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
		r.Flavors = append(r.Flavors, dto.CommonDto{Id: allFlavors[i].ID, Name: allFlavors[i].Name})
	}
	return r, nil
}

func (s *FlavorService) GetFlavor(ctx context.Context, flavorId string, authUtils *dto.AuthUtils) (*dto.GetFlavorResponse, error) {
	r := &dto.GetFlavorResponse{}
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
	getResult := flavors.Get(computeClient, flavorId)

	flavor, err := getResult.Extract()
	if err != nil {
		return nil, err
	}

	r.Flavor = append(r.Flavor, dto.FlavorDto{ID: flavor.ID, Name: flavor.Name, RAM: flavor.RAM, VCPUs: flavor.VCPUs, Disk: flavor.Disk, IsPublic: flavor.IsPublic})

	return r, nil
}
