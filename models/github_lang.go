package models

type GitHubLang struct {
	Name      string    `json:"name" validate:"required"`
	Color     string    `json:"color" validate:"required"`
	Size      int    `json:"size"`
}

type Language struct {
	Name  string
	Color string
}

type Repository struct {
	Name string
	Languages struct {
		TotalSize int
		Edges []struct {
			Size int
			Node struct {
				Language `graphql:"... on Language"`
			}
		}
	} `graphql:"languages(first: 100)"`
}

type Query struct {
	Search struct {
		Nodes []struct {
			Repository `graphql:"... on Repository"`
		}
	} `graphql:"search(first: 100, query: $q, type: REPOSITORY)"`
}
