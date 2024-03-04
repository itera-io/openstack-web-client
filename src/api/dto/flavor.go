package dto

type ListFlavorRequest struct{}

type ListFlavorResponse struct {
	Flavors []Flavor `json:"flavors"`
}

type Flavor struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}
