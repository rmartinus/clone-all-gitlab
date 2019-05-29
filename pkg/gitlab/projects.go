package gitlab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Project represents gitlab response
type Project struct {
	Namespace struct {
		FullPath string `json:"full_path"`
	}
	Path    string `json:"path_with_namespace"`
	RepoURL string `json:"http_url_to_repo"`
}

// GetProjects gets a list of projects
func GetProjects(token string, gitlabURL string, namespace string, perPage int, page int) ([]Project, error) {
	gitlabURL = fmt.Sprintf("%s?include_subgroups=true&per_page=%d&page=%d", gitlabURL, perPage, page)

	client := &http.Client{}
	req, err := http.NewRequest("GET", gitlabURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("PRIVATE-TOKEN", token)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("not OK - Status: %d. Response: %s", res.StatusCode, string(body))
	}

	var ps []Project
	err = json.Unmarshal(body, &ps)
	if err != nil {
		return nil, err
	}

	if len(namespace) > 1 {
		var filteredProjects []Project
		for _, p := range ps {
			if strings.HasPrefix(p.Namespace.FullPath, namespace) {
				filteredProjects = append(filteredProjects, p)
			}
		}
		return filteredProjects, nil
	}

	return ps, nil
}
