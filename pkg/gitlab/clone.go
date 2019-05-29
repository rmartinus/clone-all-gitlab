package gitlab

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

const errAlreadyUpToDate = "already up-to-date"

// Clone clones the given url to the current directory or do a git pull if it exists
func Clone(id int, token string, clonePath string, jobs <-chan Project, errors chan<- error) {
	for j := range jobs {
		fullPath := clonePath + j.Path
		fmt.Printf("**** Worker %d - trying to clone %s to %s\n", id, j.RepoURL, fullPath)

		if stat, err := os.Stat(fullPath); err == nil && stat.IsDir() {
			r, errPull := git.PlainOpen(fullPath)
			if errPull != nil {
				errors <- errPull
			}
			w, errPull := r.Worktree()
			if errPull != nil {
				errors <- errPull
			}
			errPull = w.Pull(&git.PullOptions{
				RemoteName: "origin",
				Auth: &http.BasicAuth{
					Password: token,
				},
				Progress: os.Stdout,
			})

			if errPull != nil && errPull.Error() != errAlreadyUpToDate {
				fmt.Printf("**** Worker %d - error pulling %s - error: %v\n", id, j.RepoURL, errPull)
				errors <- fmt.Errorf("error pulling %s - %v", j.RepoURL, errPull)
			} else if errPull != nil && errPull.Error() == errAlreadyUpToDate {
				fmt.Printf("**** Worker %d - pull %s - %s\n", id, j.RepoURL, errPull)
				errors <- nil
			} else {
				fmt.Printf("**** Worker %d - successfully pulled %s\n", id, j.RepoURL)
				errors <- nil
			}
		} else {
			_, err := git.PlainClone(fullPath, false, &git.CloneOptions{
				URL: j.RepoURL,
				Auth: &http.BasicAuth{
					Password: token,
				},
				Progress: os.Stdout,
			})

			if err != nil {
				fmt.Printf("**** Worker %d - error cloning %s - error: %v\n", id, j.RepoURL, err)
				errors <- fmt.Errorf("error cloning %s - %v", j.RepoURL, err)
			} else {
				fmt.Printf("**** Worker %d - successfully cloned %s\n", id, j.RepoURL)
				errors <- nil
			}
		}
	}
}
