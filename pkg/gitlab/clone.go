package gitlab

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

// Clone clones the given url to the current directory
func Clone(name string, url string, token string) error {
	fmt.Printf("Cloning %s to %s\n", url, name)
	_, err := git.PlainClone("/tmp/test/"+name, false, &git.CloneOptions{
		URL: url,
		Auth: &http.BasicAuth{
			Password: token,
		},
		Progress: os.Stdout,
	})

	return err
}
