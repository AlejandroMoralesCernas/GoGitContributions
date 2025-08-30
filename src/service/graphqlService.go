package service

import (
	"GoGitContributions/src/config"
	"GoGitContributions/src/data/api"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func MakeGithubGraphqlRequest(graphqlBody api.GraphqlRequest) (api.GraphQLResponse, error) {
	fmt.Println("Making GitHub GraphQL request...")

	graphqlBodyBytes, err := json.Marshal(graphqlBody)
	if err != nil {
		return api.GraphQLResponse{}, fmt.Errorf("marshaling request body: %w", err)
	}

	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(graphqlBodyBytes))
	if err != nil {
		return api.GraphQLResponse{}, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+config.GetGithubToken());
	req.Header.Set("Content-Type", "application/json");

	resp , err := client.Do(req)
	if err != nil {
		return api.GraphQLResponse{}, fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return api.GraphQLResponse{}, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var graphqlResponse api.GraphQLResponse
	if err := json.NewDecoder(resp.Body).Decode(&graphqlResponse); err != nil {
		return api.GraphQLResponse{}, fmt.Errorf("decoding response: %w", err)
	}
	
	return graphqlResponse, nil
}