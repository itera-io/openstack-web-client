package dto

type ListProjectRequest struct{}

type ListProjectResponse struct {
	Projects []CommonDto `json:"projects"`
}

type Project struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}
