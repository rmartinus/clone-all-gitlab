package main

import (
	"fmt"
	"os"

	"github.com/rmartinus/clone-all-gitlab/pkg/gitlab"
)

const (
	perPage    = 50
	workerPool = 5
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
	totalProject := 0

	jobs := make(chan gitlab.Project, 100)
	results := make(chan error, 100)

	for w := 1; w <= workerPool; w++ {
		go gitlab.Clone(w, token, jobs, results)
	}

	var errs []error
	for {
		ps, err := gitlab.GetProjects(token, url, perPage, page)
		if err != nil {
			fmt.Printf("Error retrieving projects %s\n", err)
			return
		}

		if len(ps) == 0 {
			break
		}

		fmt.Printf("Page %d, project size: %d\n", page, len(ps))
		totalProject += len(ps)

		for _, p := range ps {
			jobs <- p
		}

		for i := 0; i < len(ps); i++ {
			err := <-results
			if err != nil {
				errs = append(errs, err)
			}
		}
		page++
	}

	fmt.Println("Summary:")
	fmt.Println("Total number of projects:", totalProject)
	if len(errs) > 0 {
		fmt.Println("Total number of errors:", len(errs))
		for _, err := range errs {
			fmt.Println("-", err)
		}
	}
	close(jobs)
}
