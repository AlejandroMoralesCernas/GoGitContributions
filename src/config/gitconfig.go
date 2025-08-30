package config

import "os"

func GetGithubToken() string {
	return os.Getenv("GIT_TOKEN")
}
