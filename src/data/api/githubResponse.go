package api

type Repo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	HTMLURL  string `json:"html_url"`
}

type User struct {
	Login string `json:"login"`
	Name  string `json:"name"`
}

type GraphQLResponse struct {
	Data struct {
		User struct {
			ContributionsCollection struct {
				CommitContributionsByRepository []struct {
					Repository struct {
						Name  string `json:"name"`
						Owner struct {
							Login string `json:"login"`
						} `json:"owner"`
					} `json:"repository"`
					Contributions struct {
						TotalCount int `json:"totalCount"`
					} `json:"contributions"`
				} `json:"commitContributionsByRepository"`
			} `json:"contributionsCollection"`
		} `json:"user"`
	} `json:"data"`
}