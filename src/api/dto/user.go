package dto

import (
	"time"

	"github.com/gophercloud/gophercloud/openstack/identity/v3/users"
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

type CreateUserRequest struct {
	Name        string `json:"name" binding:"required,min=3"`
	Description string `json:"description" binding:"required,min=6"`
	// DomainID is the ID this project will belong under.
	DomainID string `json:"domain_id,omitempty"`
	// Enabled sets the project status to enabled or disabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Password is the password of the new user.
	Password string `json:"password,omitempty"`
}

type CreateUserResponse struct {
	User users.User `json:"user"`
}

type ListUserRequest struct {
	// DomainID filters the response by a domain ID.
	DomainID string `form:"domain_id"`

	// Enabled filters the response by enabled users.
	Enabled *bool `form:"enabled"`

	// Name filters the response by username.
	Name string `form:"name"`
}

type ListUserResponse struct {
	Users []users.User `json:"users"`
}
