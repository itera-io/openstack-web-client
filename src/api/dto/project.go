package dto

import "github.com/gophercloud/gophercloud/openstack/identity/v3/projects"

type ListProjectRequest struct{}

type ListProjectResponse struct {
	Projects []projects.Project `json:"projects"`
}

type Project struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type CreateProjectRequest struct {
	Name        string `json:"name" binding:"required,min=3"`
	Description string `json:"description" binding:"required,min=6"`
	// DomainID is the ID this project will belong under.
	DomainID string `json:"domain_id,omitempty"`
	// Enabled sets the project status to enabled or disabled.
	Enabled *bool `json:"enabled,omitempty"`
	// IsDomain indicates if this project is a domain.
	IsDomain *bool `json:"is_domain,omitempty"`
	// ParentID specifies the parent project of this new project.
	ParentID string `json:"parent_id,omitempty"`
	// Tags is a list of tags to associate with the project.
	Tags []string `json:"tags,omitempty"`
}
