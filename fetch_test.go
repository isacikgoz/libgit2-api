package git

import (
	"testing"
	"io/ioutil"
	"os"
)

func TestFetch(t *testing.T) {
	creds := &CredentialsAsPlainText{
	}
	dir, err := ioutil.TempDir("", "temp-fetch-dir")
	if err != nil {
		t.Fatalf("Test Failed. error: %s", err.Error())
	}
	defer os.RemoveAll(dir) // clean up
	repo, err := Clone(dir, "https://github.com/isacikgoz/gia.git", &CloneOptions{
		Credentials: creds,
	})
	if err != nil {
		t.Fatalf("Test Failed. error: %s", err.Error())
	}
	var tests = []struct {
		input *FetchOptions
		output error
	}{
		{&FetchOptions{
			Remote: "origin",
			Credentials: creds,
		}, nil},
		{&FetchOptions{
			Remote: "asda",
			Credentials: creds,
		}, ErrNotValidRemoteName},
		{&FetchOptions{
			Credentials: creds,
		}, ErrNoRemoteName},
	}
	for _, test := range tests {
		if err := repo.Fetch(test.input); err != test.output {
			t.Errorf("input: %v error: %s\n", test.input, err.Error())
		}
	}
}