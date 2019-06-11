package git

import (
	"errors"
	"path/filepath"

	lib "gopkg.in/libgit2/git2go.v27"
)

// Repository is the wrapper and main interface to git repository
type Repository struct {
	essence *lib.Repository
	path    string

	RefMap map[string][]Ref
	Head   *Branch
}

// RefType defines the ref types
type RefType uint8

// These types are used for mapping references
const (
	RefTypeTag RefType = iota
	RefTypeBranch
	RefTypeHEAD
)

// Ref is the wrapper of lib.Ref
type Ref interface {
	Type() RefType
	Target() *Commit
	String() string
}

// Open load the repository from the filesystem
func Open(path string) (*Repository, error) {
	repo, realpath, err := initRepoFromPath(path)
	if err != nil {
		return nil, ErrCannotOpenRepo
	}
	r := &Repository{
		path:    realpath,
		essence: repo,
	}
	r.RefMap = make(map[string][]Ref)
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

// Path returns the filesystem location of the repository
func (r *Repository) Path() string {
	return r.path
}
