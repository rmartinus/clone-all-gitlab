package gitlab

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Project represents gitlab response
type Project struct {
	Path    string `json:"path_with_namespace"`
	RepoURL string `json:"http_url_to_repo"`
}

// GetProjects gets a list of projects
func GetProjects(token string, url string, perPage int, page int) ([]byte, error) {
	url = fmt.Sprintf("%s?include_subgroups=true&per_page=%d&page=%d", url, perPage, page)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("new HTTP request failed with error %s", err)
	}

	req.Header.Set("PRIVATE-TOKEN", token)

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP GET failed with error %s", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("read body failed with error %s", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("not OK - Status: %d. Response: %s", res.StatusCode, string(body))
	}

	return body, nil
}
