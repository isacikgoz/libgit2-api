package main

import (
	git "github.com/isacikgoz/libgit2-api"
	"os"
	"fmt"
)

// go run fetch.go /Users/ibrahim/Development/test/sashimi
func main() {
	r, err := git.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	creds := &git.CredentialsAsSshAgent{
		UserName: "git",
	}
	opts := &git.FetchOptions{
		Remote: "origin",
		Credentials: creds,
	}
	err = r.Fetch(opts)
	if err != nil {
		panic(err)
	}
	fmt.Println("finished")
}