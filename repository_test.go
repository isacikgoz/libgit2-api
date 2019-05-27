package git

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestOpen(t *testing.T) {
	wd, _ := os.Getwd()
	var tests = []struct {
		input string
		err   error
	}{
		{"/tmp", ErrCannotOpenRepo},
		{"/", ErrCannotOpenRepo},
		{wd, nil},
	}
	for _, test := range tests {
		if _, err := Open(test.input); err != test.err {
			t.Errorf("input: %s\n error: %s", test.input, err.Error())
		}
	}
}

func testCloneFromLocal(name string) (*Repository, error) {
	wd, _ := os.Getwd()
	creds := &CredentialsAsPlainText{}
	dir, err := ioutil.TempDir("", "temp-"+name+"-dir")
	if err != nil {
		return nil, err
	}
	repo, err := Clone(dir, wd, &CloneOptions{
		Credentials: creds,
	})
	if err != nil {
		return nil, err
	}
	return repo, nil
}
