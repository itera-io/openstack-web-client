package dto

type ListProjectRequest struct{}

type ListProjectResponse struct {
	Projects []Project `json:"projects"`
}

type Project struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}
