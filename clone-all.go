package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const perPage = 20

// Project represents gitlab response
type Project struct {
	RepoURL string `json:"http_url_to_repo"`
}

func main() {
	token := os.Getenv("GITLAB_TOKEN")
	url := os.Getenv("GITLAB_URL")

	if len(token) < 1 {
		fmt.Print("Please set GITLAB_TOKEN")
	}

	if len(url) < 1 {
		fmt.Print("Please set GITLAB_URL")
	}

	page := 1

	for {
		body, err := getProjects(token, url, perPage, page)
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
			fmt.Println(p.RepoURL)
		}
		page++
	}
}

func getProjects(token string, url string, perPage int, page int) ([]byte, error) {
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
