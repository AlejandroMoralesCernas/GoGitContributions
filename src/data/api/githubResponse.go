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