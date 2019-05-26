package git

import (
	"testing"
	"io/ioutil"
	"os"
)

func TestClone(t *testing.T) {
	dirs := make([]string, 0)
	for i := 0; i < 2; i++{
		dir, err := ioutil.TempDir("", "temp-clone-dir")
		if err != nil {
			t.Fatalf("Test Failed. error: %s", err.Error())
		}
		defer os.RemoveAll(dir) // clean up
		dirs = append(dirs, dir)
	}
	creds := &CredentialsAsSshAgent{
		UserName: "git",
	}
	opts := &CloneOptions{
		Credentials: creds,
	}
	var tests = []struct {
		inputDir string
		inputURL string
		inputOpt *CloneOptions
		err error
	}{
		{dirs[0], "git@github.com:isacikgoz/libgit2-api.git", opts, nil},
		{dirs[1], "", opts, ErrClone},
	}
	for _, test := range tests {
		if _, err := Clone(test.inputDir, test.inputURL, test.inputOpt); err != test.err {
			t.Errorf("Test Failed. dir: %s, url: %s inputted, got \"%s\" as error.",
			 test.inputDir, test.inputURL, err.Error())
		}
	}
}
