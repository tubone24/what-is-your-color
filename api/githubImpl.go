package api

import (
	"context"
	"github.com/shurcooL/githubv4"
	"github.com/spf13/viper"
	"github.com/tubone24/what-is-your-color/models"
	"golang.org/x/oauth2"
)

type Language struct {
	Name  string
	Color string
}

type Repository struct {
	Name string
	Languages struct {
		TotalSize int
		Edges []struct {
			Size int
			Node struct {
				Language `graphql:"... on Language"`
			}
		}
	} `graphql:"languages(first: 100)"`
}

var query struct {
	Search struct {
		Nodes []struct {
			Repository `graphql:"... on Repository"`
		}
	} `graphql:"search(first: 100, query: $q, type: REPOSITORY)"`
}

type GithubClientImpl struct {
}

func (client *GithubClientImpl) GetColor(username string) (error, []models.GitHubLang){
	err := client.CallApi(username)
	if err != nil {
		return err, nil
	}
	var langs []models.GitHubLang
	for _, repo := range query.Search.Nodes {
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

func (client *GithubClientImpl) CallApi(username string) error {
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
		return err
	}
	return nil
}

func langsContains(arr []models.GitHubLang, str string) (bool, int){
	for i, v := range arr{
		if v.Name == str{
			return true, i
		}
	}
	return false, -1
}
