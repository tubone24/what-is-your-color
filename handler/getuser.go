package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/shurcooL/githubv4"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"

)

type GetGitHubUser struct {
	echo.Context
}

func (c *GetGitHubUser) GetUser() string{
	var query struct {
		Viewer struct {
			Login     string
		}
	}

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: viper.GetString("github.token")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)
	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		// Handle error.
	}
	return query.Viewer.Login
}
