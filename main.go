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
	opt := &github.ListOptions{Page: 10}
	
	//fmt.Println(opt)
	// get all pages of results
	for {
		repos, resp, err := client.Organizations.List(context.Background(), "CoolCaper", opt)
		fmt.Println(resp)
		fmt.Println(repos)
		//fmt.Println(client.Repositories.ListByOrg(context.Background(), "github", opt))
		if err != nil {

		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	//UNUSED VARIABLES	
}

