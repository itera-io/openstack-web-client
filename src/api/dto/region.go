package dto

type ListRegionRequest struct{}

type ListRegionResponse struct {
	Regions []CommonDto `json:"regions"`
}

type RegionDto struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}
