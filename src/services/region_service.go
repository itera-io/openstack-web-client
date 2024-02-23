package services

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/regions"
	"github.com/itera-io/openstack-web-client/api/dto"
)

type RegionService struct{}

func NewRegionService() *RegionService {
	return &RegionService{}
}

func (s *RegionService) ListRegions(req *dto.ListRegionRequest, authUtils *dto.AuthUtils) (*dto.ListRegionResponse, error) {
	r := &dto.ListRegionResponse{}
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
		return nil, err
	}
	listOpts := regions.ListOpts{}

	allPages, err := regions.List(identityClient, listOpts).AllPages()
	if err != nil {
		return nil, err
	}

	allRegions, err := regions.ExtractRegions(allPages)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(allRegions); i++ {
		r.Regions = append(r.Regions, dto.RegionItem{Id: allRegions[i].ID, Name: allRegions[i].Description})
	}
	return r, nil
}
