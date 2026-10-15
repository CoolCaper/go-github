package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v92/github"
)

func main() {
	fmt.Println("Hello, Bounty Hunter!")


	//example code from github api docs, used for  testing 

	client, err := github.NewClient()
	if err != nil {
		// Handle error.
	}
	//fmt.Println(client)
	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 0},
	}
	
	//fmt.Println(opt)
	// get all pages of results
	var allRepos []*github.Repository
	for {
		repos, resp, err := client.Repositories.ListByOrg(context.Background(), "github", opt)
		fmt.Println(client.Repositories.ListByOrg(context.Background(), "github", opt))
		if err != nil {

		}
		allRepos = append(allRepos, repos...)

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
}

