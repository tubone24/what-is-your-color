package api

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/tubone24/what-is-your-color/api/mock"
	"github.com/tubone24/what-is-your-color/models"
	"testing"
)

func TestLangContainsTrue(t *testing.T) {
	var input []models.GitHubLang
	input = append(input, models.GitHubLang{Name: "Python", Color: "test", Size: 1})
	input = append(input, models.GitHubLang{Name: "JavaScript", Color: "test2", Size: 2})
	actualBool, actualIndex := langsContains(input, "Python")
	assert.Equal(t, true, actualBool)
	assert.Equal(t, 0, actualIndex)
}

func TestLangContainsFalse(t *testing.T) {
	var input []models.GitHubLang
	input = append(input, models.GitHubLang{Name: "Python", Color: "test", Size: 1})
	input = append(input, models.GitHubLang{Name: "JavaScript", Color: "test2", Size: 2})
	actualBool, actualIndex := langsContains(input, "Go")
	assert.Equal(t, false, actualBool)
	assert.Equal(t, -1, actualIndex)
}

func TestGetColor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock.NewMockClient(ctrl)
	m.EXPECT().CallApi("tubone24").Return(nil)
	m.GetColor("tubone24")
}
