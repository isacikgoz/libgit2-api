package git

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestCommit(t *testing.T) {
	repo, err := testCloneFromLocal("commit")
	defer os.RemoveAll(repo.path) // clean up
	if err != nil {
		t.Fatalf("Test Failed. error: %s", err.Error())
	}
	if err := addTestFilesToRepo(repo); err != nil {
		t.Fatalf("Test Failed. error: %s", err.Error())
	}
	var tests = []struct {
		inputMsg string
		inputSig *Signatute
		output   error
	}{
		{"test commit", &Signatute{
			Name:  "Ibrahim Serdar Acikgoz",
			Email: "serdarcikgoz86@gmail.com",
			When:  time.Now(),
		}, nil},
	}
	for _, test := range tests {
		if _, err := repo.Commit(test.inputMsg, test.inputSig); test.output != err {
			t.Errorf("test failed. got error: %s\n", err.Error())
		}
	}
}

func TestAmend(t *testing.T) {
	repo, err := testCloneFromLocal("amend")
	defer os.RemoveAll(repo.path) // clean up
	if err != nil {
		t.Fatalf("Test Failed. error: %s", err.Error())
	}
	if err := addTestFilesToRepo(repo); err != nil {
		t.Fatalf("Test Failed. error: %s", err.Error())
	}
	sig := &Signatute{
		Name:  "Ibrahim Serdar Acikgoz",
		Email: "serdarcikgoz86@gmail.com",
		When:  time.Now(),
	}
	commit, err := repo.Commit("amaend commit testing", sig)
	if err != nil {
		t.Fatalf("Test Failed. error: %s", err.Error())
	}
	var tests = []struct {
		inputMsg string
		inputSig *Signatute
		output   error
	}{
		{commit.Message(), sig, nil},
		{"", sig, nil},
	}
	for _, test := range tests {
		if _, err := commit.Amend(test.inputMsg, test.inputSig); test.output != err {
			t.Errorf("test failed. got error: %s\n", err.Error())
		}
	}
}

func addTestFilesToRepo(repo *Repository) error {
	// create a file to add
	d1 := []byte("package git\n\nimport \"fmt\"\n\nfunc test() {\n\tfmt.Println(\"a\")\n}\n")
	err := ioutil.WriteFile(repo.path+"/added.go", d1, 0644)
	if err != nil {
		return err
	}
	// create this file to see that it is not included into commit
	err = ioutil.WriteFile(repo.path+"/untracked.go", d1, 0644)
	if err != nil {
		return err
	}
	// get the status entries
	status, err := repo.loadStatus()
	if err != nil {
		return err
	}
	// add first file "added.go" to index
	return repo.AddToIndex(status.Entities[0])
}
