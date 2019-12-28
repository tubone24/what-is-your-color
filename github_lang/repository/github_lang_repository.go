package repository

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/tubone24/what-is-your-color/github_lang"
	"github.com/tubone24/what-is-your-color/models"
)

type graphQLGitHubLanguageRepository struct {
	authToken string
}

func NewGraphQLGitHubLanguageRepository(authToken string) github_lang.Repository {
	return &graphQLGitHubLanguageRepository{authToken}
}

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
	} `graphql:"search(first: 100, query: $q, type: $searchType)"`
}

//query
//{
//  search(query: "user:tubone24", type: REPOSITORY, first: 100) {
//    pageInfo {
//      endCursor
//      startCursor
//    }
//    edges {
//      node {
//        ... on Repository {
//          name
//          isArchived
//          diskUsage
//          url
//          languages(first: 100) {
//            pageInfo {
//              endCursor
//              startCursor
//            }
//            totalCount
//            edges {
//              size
//              node {
//                name
//                color
//              }
//            }
//          }
//        }
//      }
//    }
//  }
//}

func (m *graphQLGitHubLanguageRepository) Fetch(ctx context.Context, query string, args ...interface{}) ([]*models.GitHubLang, error) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: viper.GetString("github.token")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)
	variables := map[string]interface{}{
		"q": githubql.String("GraphQL"),
		"searchType":  githubql.SearchTypeRepository,
	}
	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		// Handle error.
		fmt.Println(err)
	}
	for _, repo := range query.Search.Nodes {
		fmt.Println("---------")
		fmt.Println(repo.Name)
		for _, lang := range repo.Languages.Edges {
			fmt.Println(lang.Node.Name)
			fmt.Println(lang.Node.Color)
			fmt.Println(lang.Size)
		}
	}
}
