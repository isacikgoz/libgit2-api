package git

import (
	"testing"
	"io/ioutil"
	"os"
)

func TestClone(t *testing.T) {
	dir, err := ioutil.TempDir("", "temp-clone-dir")
	if err != nil {
		t.Fatalf("Test Failed. error: %s", err.Error())
	}
	defer os.RemoveAll(dir) // clean up
	creds := &CredentialsAsPlainText{
	}
	opts := &CloneOptions{}
	opts.Credentials = creds
	var tests = []struct {
		inputDir string
		inputURL string
		inputOpt *CloneOptions
	}{
		{dir, "https://github.com/isacikgoz/gia.git", opts},
	}
	for _, test := range tests {
		if _, err := Clone(test.inputDir, test.inputURL, test.inputOpt); err != nil {
			t.Errorf("Test Failed. dir: %s, url: %s inputted, got \"%s\" as error.",
			 test.inputDir, test.inputURL, err.Error())
		}
	}
}