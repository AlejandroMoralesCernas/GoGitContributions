package service

import (
	"fmt"
	"GoGitContributions/src/config"
	"net/http"
	"encoding/json"
	"io"
	"GoGitContributions/src/data/api"
)

func GitService() {

	fmt.Println("Git Service")
}

func GetPublicRepos() ([]api.Repo, error) {
	fmt.Println("Getting public repos...")
	url := "https://api.github.com/user/repos"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+config.GetGithubToken())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}


	var repos []api.Repo
	if err := json.Unmarshal(body, &repos); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	return repos, nil
}

