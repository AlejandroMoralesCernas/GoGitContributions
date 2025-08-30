package main
import "fmt"
import "GoGitContributions/src/service"
import "net/http"

func main() {
	mux:= http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Go Git Contributions! Health check...")
	})

	mux.HandleFunc("/repos", func(w http.ResponseWriter, r *http.Request) {
		repos, err := service.GetPublicRepos()
		if err != nil {
			http.Error(w, "Error fetching repos", http.StatusInternalServerError)
			return
		}
		for _, repo := range repos {
			fmt.Fprintf(w, "ID: %d, Name: %s, URL: %s\n", repo.ID, repo.Name, repo.HTMLURL)
		}
	})

	mux.HandleFunc("/userDetails", func(w http.ResponseWriter, r *http.Request) {
		userDetails, err := service.GetUserDetails()
		if err != nil {
			http.Error(w, "Error fetching user details", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "User Details: %s\n", userDetails)
	})

	fmt.Println("Starting Go Git Contributions on :8080...")
	http.ListenAndServe(":8080", mux)
}
