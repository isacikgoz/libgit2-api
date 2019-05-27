package git

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestMerge(t *testing.T) {
	wd, _ := os.Getwd()
	creds := &CredentialsAsPlainText{}
	dir, err := ioutil.TempDir("", "temp-merge-dir")
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
		branch string
		input  *MergeOptions
		output error
	}{
		{"asd", &MergeOptions{
			NoFF: true,
		}, ErrBranchNotFound},
		{"origin/master", &MergeOptions{
			IgnoreAlreadyUpToDate: true,
		}, nil},
	}
	for _, test := range tests {
		if err := repo.Merge(test.branch, test.input); err != test.output {
			t.Errorf("input branch: %s, error: %s\n", test.branch, err.Error())
		}
	}
}
