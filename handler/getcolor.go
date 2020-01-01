package handler

import (
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"net/http"

	"github.com/tubone24/what-is-your-color/api"
)

func GetColor() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")
		github := &api.GitHub{Client: &api.GithubClientImpl{}}
		sp := jaegertracing.CreateChildSpan(c, "Call API")
		defer sp.Finish()
		sp.SetBaggageItem("Func", "GetColor")
		sp.SetTag("Func", "GetColor")
		err, langs := github.DoGetColor(username)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal Error")
		}
		return c.JSON(http.StatusOK, langs)
	}
}
