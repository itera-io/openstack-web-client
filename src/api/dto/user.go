package dto

type ValidateUserRequest struct {
	Username               string `json:"username" binding:"required,min=3"`
	Password               string `json:"password" binding:"required,min=6"`
	Url                    string `json:"url" binding:"required"`
	Domain                 string `json:"domain"`
	ApplicationCredEnabled bool   `json:"applicationCredEnabled"`
}

type ValidateUserResponse struct {
}

type AuthenticateUserRequest struct {
	Username               string `json:"username" binding:"required,min=3"`
	Password               string `json:"password" binding:"required,min=6"`
	Url                    string `json:"url" binding:"required"`
	Domain                 string `json:"domain"`
	ApplicationCredEnabled bool   `json:"applicationCredEnabled"`
}

type AuthenticateUserResponse struct {
	Token string `json:"token"`
}
