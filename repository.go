package git

import (
	"path/filepath"
	"errors"

	lib "gopkg.in/libgit2/git2go.v27"
)

// Repository is the wrapper and main interface to git repository
type Repository struct {
	essence *lib.Repository
	path string
}

// Open load the repository from the filesystem
func Open(path string) (*Repository, error) {
	repo, realpath, err := initRepoFromPath(path)
	if err != nil {
		return nil, ErrCannotOpenRepo
	}
	r := &Repository{
		path: realpath,
		essence: repo,
	}
	return r, nil
}

func initRepoFromPath(path string) (*lib.Repository, string, error) {
	walk := path
	for {
		r, err := lib.OpenRepository(walk)
		if err == nil {
			return r, walk, err
		}
		walk = filepath.Dir(walk)
		if walk == "/" {
			break
		}
	}
	return nil, walk, errors.New("cannot load a git repository from " + path)
}