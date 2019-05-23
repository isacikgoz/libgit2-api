package git

import (
	lib "gopkg.in/libgit2/git2go.v27"
)

// Repository is the wrapper and main interface to git repository
type Repository struct {
	essence *lib.Repository
}

// CloneOptions are mostly used git clone options from a remote
type CloneOptions struct {
	Bare bool
	Recursive bool
	Depth int
}

// Clone fetches a git repository from a given url
func Clone(path string, url string, opts *CloneOptions) (*Repository, error) {
	options := &lib.CloneOptions{
		Bare: opts.Bare,
	}
	r, err:= lib.Clone(url, path, options)
	if err != nil {
		return nil, err
	}
	repository := &Repository{
		essence: r,
	}
	return repository, nil
}