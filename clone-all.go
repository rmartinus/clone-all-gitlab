package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Project represents gitlab response
type Project struct {
	RepoURL string `json:"http_url_to_repo"`
}

func main() {
	token := os.Getenv("GITLAB_TOKEN")
	url := fmt.Sprintf("%s%s", os.Getenv("GITLAB_URL"), "?include_subgroups=true&per_page=20&page=1")

	if len(token) < 1 {
		fmt.Print("Please set GITLAB_TOKEN")
	}

	if len(url) < 1 {
		fmt.Print("Please set GITLAB_URL")
	}

	body, err := getProjects(token, url)
	if err != nil {
		fmt.Printf("Error retrieving projects %s\n", err)
	}

	var ps []Project
	err = json.Unmarshal(body, &ps)
	if err != nil {
		fmt.Printf("Error unmarshalling body: %s", err)
	}

	fmt.Printf("Project size: %d\n", len(ps))

	for _, p := range ps {
		fmt.Println(p.RepoURL)
	}
}

func getProjects(token string, url string) ([]byte, error) {
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
