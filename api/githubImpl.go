package api

import (
	"context"
	"github.com/shurcooL/githubv4"
	"github.com/spf13/viper"
	"github.com/tubone24/what-is-your-color/models"
	"golang.org/x/oauth2"
)

var query = models.Query{}

type GithubClientImpl struct {
}

func (client *GithubClientImpl) GetColor(username string) (error, []models.GitHubLang){
	err, tmpQuery := client.CallApi(username)
	if err != nil {
		return err, nil
	}
	var langs []models.GitHubLang
	for _, repo := range tmpQuery.Search.Nodes {
		for _, lang := range repo.Languages.Edges {
			isContain, i := langsContains(langs, lang.Node.Name)
			if isContain {
				langs[i].Size = lang.Size + langs[i].Size
			} else {
				langs = append(langs, models.GitHubLang{Name: lang.Node.Name, Size: lang.Size, Color: lang.Node.Color})
			}
		}
	}
	return nil, langs
}

func (client *GithubClientImpl) CallApi(username string) (error, *models.Query){
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: viper.GetString("github.token")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	cc := githubv4.NewClient(httpClient)
	variables := map[string]interface{}{
		"q": githubv4.String("user:" + username),
	}
	err := cc.Query(context.Background(), &query, variables)
	if err != nil {
		return err, nil
	}
	return nil, &query
}

func langsContains(arr []models.GitHubLang, str string) (bool, int){
	for i, v := range arr{
		if v.Name == str{
			return true, i
		}
	}
	return false, -1
}
