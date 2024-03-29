package volume

import (
	"context"

	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumetypes"
	"github.com/itera-io/openstack-web-client/api/dto"
	"github.com/itera-io/openstack-web-client/pkg/logging"
)

// ListTypes lists the available volume types.
func (s *Service) ListTypes(ctx context.Context, req *dto.ListVolumeTypeRequest, authUtils *dto.AuthUtils) (*dto.ListVolumeTypeResponse, error) {
	client, err := s.newBlockStorageClient(authUtils)
	if err != nil {
		return nil, err
	}

	listOpts := volumetypes.ListOpts{}
	allPages, err := volumetypes.List(client, listOpts).AllPages()
	if err != nil {
		s.Logger.Error(logging.IdentityClient, logging.ExternalService, "Failed to list volume types", nil)
		return nil, err
	}

	allVolumeTypes, err := volumetypes.ExtractVolumeTypes(allPages)
	if err != nil {
		return nil, err
	}

	return &dto.ListVolumeTypeResponse{VolumeTypes: allVolumeTypes}, nil
}
