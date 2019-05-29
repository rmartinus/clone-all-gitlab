package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rmartinus/clone-all-gitlab/pkg/gitlab"
)

const (
	perPage          = 50
	workerPool       = 5
	gitlabURL        = "https://gitlab.com/api/v4/groups/your-org/projects/"
	defaultClonePath = "/tmp/clone-all/"
)

func main() {
	fmt.Println(gitlab.Banner)

	token := os.Getenv("GITLAB_TOKEN")
	clonePath := os.Getenv("DEV_HOME")
	namespace := os.Getenv("GITLAB_NAMESPACE")

	if len(token) < 1 {
		fmt.Println("GITLAB_TOKEN is not set")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	if len(clonePath) == 0 {
		text := readLine("Clone path (defaults to /tmp/clone-all/): ", scanner)

		if len(text) > 0 {
			clonePath = text
		} else {
			clonePath = defaultClonePath
		}
	}

	if clonePath[len(clonePath)-1:] != "/" {
		clonePath += "/"
	}

	if len(namespace) == 0 {
		text := readLine("Namespace (eg. <your-org-name>/<repository>): ", scanner)
		if len(text) > 0 {
			namespace = text
		}
	}
	cloneProjects(token, clonePath, namespace)
}

func readLine(s string, scanner *bufio.Scanner) string {
	fmt.Print(s)
	scanner.Scan()
	input := scanner.Text()
	return input
}

func cloneProjects(token string, clonePath string, namespace string) {
	page := 1
	totalProject := 0

	jobs := make(chan gitlab.Project, 100)
	errors := make(chan error, 100)

	for w := 1; w <= workerPool; w++ {
		go gitlab.Clone(w, token, clonePath, jobs, errors)
	}

	var errs []error
	for {
		ps, err := gitlab.GetProjects(token, gitlabURL, namespace, perPage, page)
		if err != nil {
			fmt.Printf("Error retrieving projects %s\n", err)
			return
		}

		if len(ps) == 0 {
			break
		}

		totalProject += len(ps)

		for _, p := range ps {
			jobs <- p
		}

		for i := 0; i < len(ps); i++ {
			err := <-errors
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
