package gitlab

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

// Clone clones the given url to the current directory
func Clone(id int, token string, jobs <-chan Project, results chan<- bool) {
	for j := range jobs {
		fmt.Printf("**** Worker %d - cloning %s to %s\n", id, j.RepoURL, j.Path)
		_, err := git.PlainClone("/tmp/clone-all/"+j.Path, false, &git.CloneOptions{
			URL: j.RepoURL,
			Auth: &http.BasicAuth{
				Password: token,
			},
			Progress: os.Stdout,
		})

		if err != nil {
			fmt.Printf("**** Worker %d - error cloning %s - error: %v\n", id, j.RepoURL, err)
			results <- false
		} else {
			fmt.Printf("**** Worker %d - successfully cloned %s\n", id, j.RepoURL)
			results <- true
		}
	}
}
