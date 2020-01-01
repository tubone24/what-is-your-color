package handler

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/labstack/echo/v4"
)

const userJSON = `{"name":"Jon Snow","email":"jon@labstack.com"}`

func TestGetColor(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/get/:username")
	c.SetParamNames("username")
	c.SetParamValues("tubone24")
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, userJSON, rec.Body.String())
}

