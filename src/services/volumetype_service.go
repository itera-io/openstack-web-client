package services

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumetypes"
	"github.com/itera-io/openstack-web-client/api/dto"
	"github.com/itera-io/openstack-web-client/config"
	"github.com/itera-io/openstack-web-client/pkg/logging"
)

type VolumeTypeService struct{ Logger logging.Logger }

func NewVolumeTypeService(cfg *config.Config) *VolumeTypeService {
	return &VolumeTypeService{Logger: logging.NewLogger(cfg)}
}

func (s *VolumeTypeService) List(ctx context.Context, req *dto.ListVolumeTypeRequest, authUtils *dto.AuthUtils) (*dto.ListVolumeTypeResponse, error) {
	r := &dto.ListVolumeTypeResponse{}
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: authUtils.BaseUrl,
		TokenID:          authUtils.Token,
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		return nil, err
	}
	client, err := openstack.NewBlockStorageV3(provider, gophercloud.EndpointOpts{})
	if err != nil {
		s.Logger.Error(logging.IdentityClient, logging.ExternalService, err.Error(), nil)
		return nil, err
	}
	listOpts := volumetypes.ListOpts{}
	allPages, err := volumetypes.List(client, listOpts).AllPages()
	if err != nil {
		return nil, err
	}

	allVolumeTypes, err := volumetypes.ExtractVolumeTypes(allPages)
	if err != nil {
		return nil, err
	}
	r.VolumeTypes = allVolumeTypes
	return r, nil
}
