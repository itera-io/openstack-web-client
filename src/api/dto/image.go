package dto

type ListImageRequest struct {
}

type ListImageResponse struct {
	Images []CommonDto `json:"images"`
}
type GetImageResponse struct {
	Image []ImageDto `json:"image"`
}

type ImageDto struct {
	// ID is the flavor's unique ID.
	ID string `json:"id"`

	// Disk is the amount of root disk, measured in GB.
	MinDisk int `json:"disk"`

	// RAM is the amount of memory, measured in MB.
	MinRAM int `json:"ram"`

	// Name is the name of the flavor.
	Name string `json:"name"`

	// VCPUs indicates how many (virtual) CPUs are available for this flavor.
	Status string `json:"status"`
}
