package main
import("fmt"; "github.com/google/go-github/v92/github"; "context")

func main() {
	fmt.Println("Hello, Bounty Hunter!")

	test, try := github.NewClient()
	if try != nil {
	
	}
	var repos = &github.RepositoryListByOrgOptions{Type: "public"}
	fmt.Println(orgs)
}

