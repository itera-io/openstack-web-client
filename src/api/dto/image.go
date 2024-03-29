package dto

type ListImageRequest struct {
	// Owner filters on the project ID of the image.
	Owner string `form:"owner"`

	// Tags filters on specific image tags.
	Tags []string `form:"tag"`
	// Status filters on the status of the image.
	// Multiple statuses can be specified by constructing a string
	// such as "in:saving,queued,active".
	Status string `form:"status"`
}

type ListImageResponse struct {
	Images []ImageListItem `json:"images"`
}
type GetImageResponse struct {
	Image []ImageDto `json:"image"`
}

type ImageDto struct {
	// ID is the flavor's unique ID.
	ID string `json:"id"`

	// Disk is the amount of root disk, measured in GB.
	MinDiskGigabytes int `json:"disk"`

	// RAM is the amount of memory, measured in MB.
	MinRAMMegabytes int `json:"ram"`

	// Name is the name of the flavor.
	Name string `json:"name"`

	// VCPUs indicates how many (virtual) CPUs are available for this flavor.
	Status string `json:"status"`
}

type ImageListItem struct {
	// ID is the flavor's unique ID.
	ID string `json:"id"`

	// Disk is the amount of root disk, measured in GB.
	MinDisk int `json:"disk"`

	// RAM is the amount of memory, measured in MB.
	MinRAM int `json:"ram"`

	// Name is the name of the flavor.
	Name string `json:"name"`

	// Tags filters on specific image tags.
	Tags []string `json:"tag"`
}
