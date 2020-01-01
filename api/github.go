package api

import "github.com/tubone24/what-is-your-color/models"

type GitHub struct {
	Client Client
}

type Client interface {
	GetColor(username string) (error, []models.GitHubLang)
}

func (github *GitHub) DoGetColor(username string) (error, []models.GitHubLang) {
	return github.Client.GetColor(username)
}
