package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	//"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"

	"github.com/tubone24/what-is-your-color/middleware"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

type GetGitHubLang struct {
	echo.Context
}

func (c *GetGitHubLang) Foo() string{
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
	println("aaa")
	return query.Viewer.Login
}

func main() {
	e := echo.New()

	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)

	//timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &GetGitHubLang{c}
			return next(cc)
		}
	})

	e.GET("/", func(c echo.Context) error {
		cc := c.(*GetGitHubLang)
		aaa := cc.Foo()
		return cc.JSON(http.StatusOK, map[string]string{"Login": aaa})
	})

	log.Fatal(e.Start(viper.GetString("server.address")))
}
