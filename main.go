package main
import("fmt"; "github.com/google/go-github/v92/github"; "context")

func main() {
	fmt.Println("Hello, Bounty Hunter!")

	test, try := github.NewClient()
	if try != nil {
	
	}

	o, _, try := test.Organizations.List(context.Background(), "CoolCaper", nil)
	fmt.Println(o)
}

