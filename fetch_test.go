package git

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFetch(t *testing.T) {
	wd, _ := os.Getwd()
	creds := &CredentialsAsPlainText{}
	dir, err := ioutil.TempDir("", "temp-fetch-dir")
	if err != nil {
		t.Fatalf("Test Failed. error: %s", err.Error())
	}
	defer os.RemoveAll(dir) // clean up
	repo, err := Clone(dir, wd, &CloneOptions{
		Credentials: creds,
	})
	if err != nil {
		t.Fatalf("Test Failed. error: %s", err.Error())
	}
	var tests = []struct {
		input  *FetchOptions
		output error
	}{
		{&FetchOptions{
			Remote:      "origin",
			Tags:        true,
			Credentials: creds,
		}, nil},
		{&FetchOptions{
			Remote:      "asda",
			Credentials: creds,
		}, ErrNotValidRemoteName},
		{&FetchOptions{
			Credentials: creds,
			Prune:       true,
		}, ErrNoRemoteName},
	}
	for _, test := range tests {
		if err := repo.Fetch(test.input); err != test.output {
			t.Errorf("input: %v error: %s\n", test.input, err.Error())
		}
	}
}
