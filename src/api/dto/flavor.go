package dto

type ListFlavorRequest struct{}

type ListFlavorResponse struct {
	Flavors []CommonDto `json:"flavors"`
}
type GetFlavorResponse struct {
	Flavor []FlavorDto `json:"flavor"`
}

type FlavorDto struct {
	// ID is the flavor's unique ID.
	ID string `json:"id"`

	// Disk is the amount of root disk, measured in GB.
	Disk int `json:"disk"`

	// RAM is the amount of memory, measured in MB.
	RAM int `json:"ram"`

	// Name is the name of the flavor.
	Name string `json:"name"`

	// VCPUs indicates how many (virtual) CPUs are available for this flavor.
	VCPUs int `json:"vcpus"`

	// IsPublic indicates whether the flavor is public.
	IsPublic bool `json:"os-flavor-access:is_public"`
}
