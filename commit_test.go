package git

import (
	"os"
	"testing"
	"time"
)

func TestCommit(t *testing.T) {
	repo, err := testCloneFromLocal("commit")
	defer os.RemoveAll(repo.path) // clean up
	checkFatal(t, err)

	err = addTestFilesToRepo(repo)
	checkFatal(t, err)

	var tests = []struct {
		inputMsg string
		inputSig *Signatute
		output   error
	}{
		{"test commit", &Signatute{
			Name:  "Some Guy",
			Email: "guysome@gmail.com",
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
	checkFatal(t, err)

	err = addTestFilesToRepo(repo)
	checkFatal(t, err)

	sig := &Signatute{
		Name:  "Some Guy",
		Email: "guysome@gmail.com",
		When:  time.Now(),
	}
	commit, err := repo.Commit("amaend commit testing", sig)
	checkFatal(t, err)

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
	if err := writeTestFile(repo.path + "/added.go"); err != nil {
		return err
	}

	// create this file to see that it is not included into commit
	if err := writeTestFile(repo.path + "/untracked.go"); err != nil {
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
