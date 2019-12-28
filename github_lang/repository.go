package github_lang

import (
	"context"

	"github.com/tubone24/what-is-your-color/models"
)

// Repository represent the article's repository contract
type Repository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []*models.GitHubLang, nextCursor string, err error)
	GetByName(ctx context.Context, id int64) (*models.GitHubLang, error)
	GetByColor(ctx context.Context, title string) (*models.GitHubLang, error)
}
