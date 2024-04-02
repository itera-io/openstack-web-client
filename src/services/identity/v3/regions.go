package v3

import (
	"context"

	"github.com/gophercloud/gophercloud/openstack/identity/v3/regions"
	"github.com/itera-io/openstack-web-client/api/dto"
)

func (s *Service) ListRegions(ctx context.Context, req *dto.ListRegionRequest, authUtils *dto.AuthUtils) (*dto.ListRegionResponse, error) {
	r := &dto.ListRegionResponse{}
	client, err := s.newIdetityV3Client(authUtils)
	if err != nil {
		return nil, err
	}
	listOpts := regions.ListOpts{}

	allPages, err := regions.List(client, listOpts).AllPages()
	if err != nil {
		return nil, err
	}

	allRegions, err := regions.ExtractRegions(allPages)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(allRegions); i++ {
		r.Regions = append(r.Regions, dto.CommonDto{Id: allRegions[i].ID, Name: allRegions[i].Description})
	}
	return r, nil
}
