package models

type GitHubLang struct {
	Name      string    `json:"name" validate:"required"`
	Color     string    `json:"color" validate:"required"`
	Size      int    `json:"size"`
}
