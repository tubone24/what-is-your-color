package api

import (
	"github.com/tubone24/what-is-your-color/models"
	"testing"
	"github.com/stretchr/testify/assert"
)

type TestClientImpl struct {
}

func TestlangContainsTrue(t *testing.T) {
	var input []models.GitHubLang
	input = append(input, models.GitHubLang{Name: "Python", Color: "test", Size: 1})
	input = append(input, models.GitHubLang{Name: "JavaScript", Color: "test2", Size: 2})
	actualBool, actualIndex := langsContains(input, "Python")
	assert.Equal(t, true, actualBool)
	assert.Equal(t, 0, actualIndex)
}