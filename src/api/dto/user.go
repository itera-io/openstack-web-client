package dto

import (
	"time"
)

type ValidateUserRequest struct {
	Username string `json:"username" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=6"`
	Url      string `json:"url" binding:"required"`
	Domain   string `json:"domain"`
}

type ValidateUserResponse struct {
}

type AuthenticateUserRequest struct {
	Username string `json:"username" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=6"`
	Url      string `json:"url" binding:"required"`
	Domain   string `json:"domain"`
}

type AuthenticateUserResponse struct {
	Token   Token  `json:"token"`
	TokenID string `json:"tokenId"`
}

type Domain struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type ProjectDto struct {
	Domain Domain `json:"domain"`
	ID     string `json:"id"`
	Name   string `json:"name"`
}
type Role struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type User struct {
	Domain            Domain `json:"domain"`
	ID                string `json:"id"`
	Name              string `json:"name"`
	PasswordExpiresAt any    `json:"password_expires_at"`
}
type Token struct {
	ExpiresAt time.Time  `json:"expires_at"`
	IsDomain  bool       `json:"is_domain"`
	Project   ProjectDto `json:"project"`
	User      User       `json:"user"`
	Roles     []Role     `json:"roles"`
}
