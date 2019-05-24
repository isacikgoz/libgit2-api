package main

import (
	git "github.com/isacikgoz/libgit2-api"
	"os"
	"fmt"
)

// go run clone.go /Users/ibrahim/Development/sig git@github.com:isacikgoz/sig.git
func main() {
	// https://github.com/libgit2/libgit2/issues/3392 as implied here, github uses git as username
	creds := &git.CredentialsAsSshAgent{
		UserName: "git",
	}
	_, err := git.Clone(os.Args[1], os.Args[2], &git.CloneOptions{
		Bare: false,
		Credentials: creds,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("finished")
}