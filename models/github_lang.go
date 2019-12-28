package models


// GitHubLang represent the article model
type GitHubLang struct {
	Name      string    `json:"name" validate:"required"`
	Color     string    `json:"color" validate:"required"`
	Size      string    `json:"size"`
}
