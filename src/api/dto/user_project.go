package dto

type ListUserProjectRequest struct{}

type ListUserProjectResponse struct {
	UserProjects []UserProject `json:"projects"`
}

type UserProject struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}
