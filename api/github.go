package api

import "github.com/tubone24/what-is-your-color/models"

type GitHub struct {
	Client Client
}

type Client interface {
	GetColor(username string) (error, []models.GitHubLang)
	CallApi(username string) error
}

func (github *GitHub) DoGetColor(username string) (error, []models.GitHubLang) {
	return github.Client.GetColor(username)
}

func (github *GitHub) DoCallApi(username string) error {
	return github.Client.CallApi(username)
}
