package dto

type AuthenticateUserRequest struct {
	Username               string `json:"username" binding:"required,min=3"`
	Password               string `json:"password" binding:"required,password,min=6"`
	Url                    string `json:"url" binding:"required,password,min=6"`
	Domain                 string `json:"domain"`
	ApplicationCredEnabled bool   `json:"applicationCredEnabled"`
}

type ValidateUserResponse struct {
}
