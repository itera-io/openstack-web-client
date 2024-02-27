package dto

type ListRegionRequest struct{}

type ListRegionResponse struct {
	Regions []RegionItem `json:"regions"`
}

type RegionItem struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}
