package main

import (
	git "github.com/isacikgoz/libgit2-api"
	"os"
	"fmt"
)

// go run clone.go /Users/ibrahim/Development/sig https://github.com/isacikgoz/sig.git
func main() {
	_, err := git.Clone(os.Args[1], os.Args[2], &git.CloneOptions{
		Bare: false,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("finished")
}