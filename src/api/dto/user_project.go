package dto

type ListUserProjectRequest struct{}

type ListUserProjectResponse struct {
	UserProjects []CommonDto `json:"projects"`
}
