package dto

type ListRegionRequest struct{}

type ListRegionResponse struct {
	Regions []CommonDto `json:"regions"`
}
