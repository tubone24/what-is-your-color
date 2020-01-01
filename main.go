package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"

	//"time"

	"github.com/labstack/echo/v4"
	"github.com/shurcooL/githubv4"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"

	myMdl "github.com/tubone24/what-is-your-color/middleware"
	"github.com/tubone24/what-is-your-color/handler"
	"github.com/labstack/echo-contrib/jaegertracing"
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
	return query.Viewer.Login
}

func main() {
	e := echo.New()

	middL := myMdl.InitMiddleware()
	e.Use(middL.CORS)
	e.Use(middleware.Logger())
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:  1 << 10, // 1 KB
	}))
	e.HTTPErrorHandler = handler.CustomHTTPErrorHandler
	c := jaegertracing.New(e, nil)
	defer c.Close()

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
	e.GET("/get/:username", handler.GetColor())

	log.Fatal(e.Start(viper.GetString("server.address")))
}
