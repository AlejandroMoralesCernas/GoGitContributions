package service

import (
	"fmt"
	"GoGitContributions/src/config"
	"net/http"
	"encoding/json"
	"GoGitContributions/src/data/api"
	"time"
)

func GetPublicRepos() ([]api.Repo, error) {
	fmt.Println("Getting public repos...")

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", "https://api.github.com/user/repos", nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+config.GetGithubToken())

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var repos []api.Repo
	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	return repos, nil
}


func GetUserDetails() (string, error) {
	fmt.Println("Getting user details...")

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return "", fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+config.GetGithubToken())

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var user api.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return "", fmt.Errorf("decoding response: %w", err)
	}

	return user.Login, nil
}

func MakeGithubContributionRequest() (api.GraphQLResponse, error) {
	fmt.Println("Making GitHub GraphQL request...")

	query := `
	query($username: String!, $from: DateTime!, $to: DateTime!) {
		user(login: $username) {
			contributionsCollection(from: $from, to: $to) {
				commitContributionsByRepository {
					repository { name owner { login } }
					contributions(first: 100) { totalCount }
				}
				pullRequestContributionsByRepository {
					repository { name owner { login } }
					contributions(first: 100) { totalCount }
				}
				issueContributionsByRepository {
					repository { name owner { login } }
					contributions(first: 100) { totalCount }
				}
				pullRequestReviewContributionsByRepository {
					repository { name owner { login } }
					contributions(first: 100) { totalCount }
				}
			}
		}
	}`

	userDetails, err := GetUserDetails()
	if err != nil {
		return api.GraphQLResponse{}, fmt.Errorf("getting user details: %w", err)
	}

	toTime := time.Now()
	fromTime := toTime.AddDate(0, -2, 0)
	
	requestBody := api.GraphqlRequest{
		Query: query,
		Variables: map[string]interface{}{
			"username": userDetails,
			"from":     fromTime.Format(time.RFC3339),
			"to":       toTime.Format(time.RFC3339),
		},
	}
	return MakeGithubGraphqlRequest(requestBody)
}
/* func MakeGithubGraphQLRequest() string {
	client := &http.Client{Timeout: 10 * time.Second}

	// Build request body as map
	body := map[string]interface{}{
		"query": query,
		"variables": map[string]interface{}{
			"username": "octocat",
			"from":     "2025-01-01T00:00:00Z",
			"to":       "2025-12-31T23:59:59Z",
		},
	}
	
	// Encode request body directly to a buffer
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "https://api.github.com/graphql", &buf)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+config.GetGithubToken())
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("unexpected status: %s", resp.Status)
	}
	//var result map[string]interface{}


} */