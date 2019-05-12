package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rmartinus/clone-all-gitlab/pkg/gitlab"
)

const (
	perPage    = 20
	workerPool = 3
)

func main() {
	token := os.Getenv("GITLAB_TOKEN")
	url := os.Getenv("GITLAB_URL")

	if len(token) < 1 {
		fmt.Println("GITLAB_TOKEN is not set")
		return
	}

	if len(url) < 1 {
		fmt.Println("GITLAB_URL is not set")
		return
	}

	page := 1

	jobs := make(chan gitlab.Project, 100)
	results := make(chan bool, 100)

	for w := 1; w <= workerPool; w++ {
		go gitlab.Clone(w, token, jobs, results)
	}

	// for {
	body, err := gitlab.GetProjects(token, url, perPage, page)
	if err != nil {
		fmt.Printf("Error retrieving projects %s\n", err)
	}

	var ps []gitlab.Project
	err = json.Unmarshal(body, &ps)
	if err != nil {
		fmt.Printf("Error unmarshalling body: %s", err)
	}

	// if len(ps) == 0 {
	// 	break
	// }

	fmt.Printf("Page %d, project size: %d\n", page, len(ps))

	for _, p := range ps {
		jobs <- p
	}
	close(jobs)

	for i := 0; i < len(ps); i++ {
		<-results
	}

	// page++
	// }
}
