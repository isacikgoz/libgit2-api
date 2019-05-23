package main

import (
	git "github.com/isacikgoz/libgit2-api"
	"fmt"
)

func main() {
	_, err := git.Clone("/Users/ibrahim/Development/gia", "https://github.com/isacikgoz/gia.git", &git.CloneOptions{
		Bare: false,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("finished")
}