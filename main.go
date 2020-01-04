package main

import (
	"fmt"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"

	//"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/tubone24/what-is-your-color/handler"
	myMdl "github.com/tubone24/what-is-your-color/middleware"
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
			cc := &handler.GetGitHubUser{c}
			return next(cc)
		}
	})

	e.GET("/", func(c echo.Context) error {
		cc := c.(*handler.GetGitHubUser)
		aaa := cc.GetUser()
		return cc.JSON(http.StatusOK, map[string]string{"Login": aaa})
	})
	e.GET("/get/:username", handler.GetColor())

	log.Fatal(e.Start(viper.GetString("server.address")))
}
