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
	IsAuthenticated bool `json:"isAuthenticated"`
}

type AuthenticateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Url      string `json:"url" binding:"required"`
	Domain   string `json:"domain"`
	// Authentication through Application Credentials requires supplying name, project and secret
	// For project we can use TenantID
	ApplicationCredentialID     string `json:"applicationCredentialID"`
	ApplicationCredentialName   string `json:"applicationCredentialName"`
	ApplicationCredentialSecret string `json:"applicationCredentialSecret"`
}

type AuthenticateAdminUserRequest struct {
	Url    string `json:"url" binding:"required"`
	Domain string `json:"domain"`
	// Authentication through Application Credentials requires supplying name, project and secret
	// For project we can use TenantID
	ApplicationCredentialID     string `json:"-"`
	ApplicationCredentialName   string `json:"-"`
	ApplicationCredentialSecret string `json:"-"`
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
