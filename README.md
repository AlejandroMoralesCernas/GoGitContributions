# GoGitContributions

GoGitContributions is a Go web service that fetches public GitHub repositories and user details using the GitHub API.

## Features (subject to change)
- REST API endpoints for:
  - Health check (`/health`)
  - List public repositories (`/repos`)
  - Get user details (`/userDetails`)

## Prerequisites
- Download [Docker](https://www.docker.com/)
- A GitHub personal access token

## Setup
1. **Clone the repository:**
	```sh
	git clone <repo-url>
	cd GoGitContributions
	```

2. **Configure environment variables:**
	- Copy `.env.example` to `.env`:
	  ```sh
	  cp .env.example .env
	  ```
	- Edit `.env` and set your GitHub token:
	  ```env
	  GIT_TOKEN=your_github_token_here
	  ```

## Running the Project

### Using Docker Compose
1. Build and start the service:
	```sh
	docker compose up -d --build
	```
2. The API will be available at `http://localhost:8080`

## Stopping the Service
```sh
docker compose down
```

---
For any issues, please open an issue or pull request.
