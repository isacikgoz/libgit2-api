package git

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestAddToIndex(t *testing.T) {
	repo, err := testCloneFromLocal("add")
	defer os.RemoveAll(repo.path) // clean up
	status, err := repo.loadStatus()
	if err != nil {
		t.Fatalf("Test Failed. error: %s", err.Error())
	}
	if err != nil {
		t.Fatalf("Test Failed. error: %s", err.Error())
	}
	// create a file to add
	d1 := []byte("package git\n\nimport \"fmt\"\n\nfunc test() {\n\tfmt.Println(\"a\")\n}\n")
	if err := ioutil.WriteFile(repo.path+"/added.go", d1, 0644); err != nil {
		t.Fatalf("Test Failed. error: %s", err.Error())
	}
	// get the status entries
	status, err = repo.loadStatus()
	if err != nil {
		t.Fatalf("Test Failed. error: %s", err.Error())
	}
	var tests = []struct {
		input  *StatusEntry
		output error
	}{
		{status.Entities[0], nil},
	}
	for _, test := range tests {
		if err := repo.AddToIndex(test.input); err != test.output {
			t.Errorf("input: %s, output: %s\n", test.input.diffDelta.OldFile.Path, err.Error())
		}
	}
}

func TestRemoveFromIndex(t *testing.T) {
	repo, err := testCloneFromLocal("reset")
	defer os.RemoveAll(repo.path) // clean up
	if err != nil {
		t.Fatalf("Test Failed. error: %s", err.Error())
	}
	// create a file to add
	d1 := []byte("package git\n\nimport \"fmt\"\n\nfunc test() {\n\tfmt.Println(\"a\")\n}\n")
	if err := ioutil.WriteFile(repo.path+"/added.go", d1, 0644); err != nil {
		t.Fatalf("Test Failed. error: %s", err.Error())
	}
	// create this file to see that it is not included into commit
	if err := ioutil.WriteFile(repo.path+"/untracked.go", d1, 0644); err != nil {
		t.Fatalf("Test Failed. error: %s", err.Error())
	}
	// get the status entries
	status, err := repo.loadStatus()
	if err != nil {
		t.Fatalf("Test Failed. error: %s", err.Error())
	}
	if err := repo.AddToIndex(status.Entities[0]); err != nil {
		t.Fatalf("Test Failed. error: %s", err.Error())
	}
	// reload status to get new file stats
	status, err = repo.loadStatus()
	if err != nil {
		t.Fatalf("Test Failed. error: %s", err.Error())
	}
	var tests = []struct {
		input  *StatusEntry
		output error
	}{
		{status.Entities[0], nil},
		{status.Entities[1], ErrEntryNotIndexed},
	}
	for _, test := range tests {
		if err := repo.RemoveFromIndex(test.input); err != test.output {
			t.Errorf("input: %s, output: %s\n", test.input.diffDelta.OldFile.Path, err.Error())
		}
	}
}
