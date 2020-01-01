package handler

import (
	"github.com/stretchr/testify/assert"
	"github.com/tubone24/what-is-your-color/api"
	"github.com/tubone24/what-is-your-color/models"

	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/labstack/echo/v4"
)

type testGitHubClientImpl struct {
}

func (client *testGitHubClientImpl) GetColor(username string) (error, []models.GitHubLang){
	return nil, []models.GitHubLang{models.GitHubLang{Name: "Python", Color: "test", Size: 1}}
}

const userJSON = ``

func TestGetColor(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/get/:username")
	c.SetParamNames("username")
	c.SetParamValues("tubone24")
	github := &api.GitHub{Client: &testGitHubClientImpl{}}
	_ = github.Client.GetColor
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, userJSON, rec.Body.String())
}

