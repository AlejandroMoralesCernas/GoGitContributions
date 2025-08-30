package api

type GithubContributionGraphql struct {

}

type GraphqlRequest struct {
	Query string `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}
