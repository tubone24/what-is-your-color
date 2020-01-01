package handler

import (
	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/labstack/echo/v4"
)

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
	//h := GetColor()
	//require.NoError(t, h(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, userJSON, rec.Body.String())
}

