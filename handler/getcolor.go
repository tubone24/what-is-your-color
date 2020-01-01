package handler

import (
	"context"
	"fmt"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"net/http"

	"github.com/spf13/viper"
	"github.com/tubone24/what-is-your-color/models"
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

func GetColor() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")
		src := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: viper.GetString("github.token")},
		)
		httpClient := oauth2.NewClient(context.Background(), src)

		client := githubv4.NewClient(httpClient)
		variables := map[string]interface{}{
			"q": githubv4.String("user:" + username),
		}
		err := client.Query(context.Background(), &query, variables)
		sp := jaegertracing.CreateChildSpan(c, "Access GitHub")
		defer sp.Finish()
		sp.SetBaggageItem("API Access", "GitHub v4")
		sp.SetTag("Func", "GitHub v4")
		if err != nil {
			// Handle error.
			fmt.Println(err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal Error")
		}
		//langs := map[string]int{}
		var langs []models.GitHubLang
		for _, repo := range query.Search.Nodes {
			fmt.Println("---------")
			fmt.Println(repo.Name)
			for _, lang := range repo.Languages.Edges {
				isContain, i := langsContains(langs, lang.Node.Name)
				if isContain {
					langs[i].Size = lang.Size + langs[i].Size
				} else {
					langs = append(langs, models.GitHubLang{Name: lang.Node.Name, Size: lang.Size, Color: lang.Node.Color})
				}
			}
		}
		sp2 := jaegertracing.CreateChildSpan(c, "Create Returns")
		defer sp2.Finish()
		sp2.SetBaggageItem("Create Returns", "Create Returns")
		sp2.SetTag("Func", "Create Returns")
		return c.JSON(http.StatusOK, langs)
	}
}

func langsContains(arr []models.GitHubLang, str string) (bool, int){
	for i, v := range arr{
		if v.Name == str{
			return true, i
		}
	}
	return false, -1
}
