package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rmartinus/clone-all-gitlab/pkg/gitlab"
)

const perPage = 20

// Project represents gitlab response
type Project struct {
	Name    string `json:"name"`
	RepoURL string `json:"http_url_to_repo"`
}

func main() {
	token := os.Getenv("GITLAB_TOKEN")
	url := os.Getenv("GITLAB_URL")

	if len(token) < 1 {
		fmt.Print("GITLAB_TOKEN is not set")
	}

	if len(url) < 1 {
		fmt.Print("GITLAB_URL is not set")
	}

	page := 1

	for {
		body, err := gitlab.GetProjects(token, url, perPage, page)
		if err != nil {
			fmt.Printf("Error retrieving projects %s\n", err)
		}

		var ps []Project
		err = json.Unmarshal(body, &ps)
		if err != nil {
			fmt.Printf("Error unmarshalling body: %s", err)
		}

		if len(ps) == 0 {
			break
		}

		fmt.Printf("Page %d, project size: %d\n", page, len(ps))

		for _, p := range ps {
			err = gitlab.Clone(p.Name, p.RepoURL, token)
			if err != nil {
				fmt.Printf("Error cloning %s - error: %v\n", url, err)
			}
		}
		page++
	}
}
