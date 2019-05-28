package git

import (
	"time"

	lib "gopkg.in/libgit2/git2go.v27"
)

// Commit is the wrapper of actual lib.Commit object
type Commit struct {
	essence *lib.Commit
	owner   *Repository
}

// Signatute is the person who signs a commit
type Signatute struct {
	Name  string
	Email string
	When  time.Time
}

func (s *Signatute) toNewLibSignature() *lib.Signature {
	return &lib.Signature{
		Name:  s.Name,
		Email: s.Email,
		When:  s.When,
	}
}

// Commit adds a new commit onject to repository
// warning: this function does not check if the changes are indexed
func (r *Repository) Commit(message string, author ...*Signatute) (*Commit, error) {
	repo := r.essence
	head, err := repo.Head()
	if err != nil {
		return nil, err
	}
	defer head.Free()
	parent, err := repo.LookupCommit(head.Target())
	if err != nil {
		return nil, err
	}
	defer parent.Free()
	index, err := repo.Index()
	if err != nil {
		return nil, err
	}
	defer index.Free()
	treeid, err := index.WriteTree()
	if err != nil {
		return nil, err
	}
	tree, err := repo.LookupTree(treeid)
	if err != nil {
		return nil, err
	}
	defer tree.Free()
	oid, err := repo.CreateCommit("HEAD", author[0].toNewLibSignature(), author[0].toNewLibSignature(), message, tree, parent)
	if err != nil {
		return nil, err
	}
	commit, err := repo.LookupCommit(oid)
	if err != nil {
		return nil, err
	}
	return &Commit{
		essence: commit,
		owner:   r,
	}, nil
}

// Amend updates the commit and returns NEW commit pointer
func (c *Commit) Amend(message string, author ...*Signatute) (*Commit, error) {
	repo := c.owner.essence
	index, err := repo.Index()
	if err != nil {
		return nil, err
	}
	defer index.Free()
	treeid, err := index.WriteTree()
	if err != nil {
		return nil, err
	}
	tree, err := repo.LookupTree(treeid)
	if err != nil {
		return nil, err
	}
	defer tree.Free()
	// var msg string
	// if len(message) <= 0 {

	// }
	oid, err := c.essence.Amend("HEAD", author[0].toNewLibSignature(), author[0].toNewLibSignature(), message, tree)
	if err != nil {
		return nil, err
	}
	commit, err := repo.LookupCommit(oid)
	if err != nil {
		return nil, err
	}
	return &Commit{
		essence: commit,
		owner:   c.owner,
	}, nil
}

// Message returns the commit message
func (c *Commit) Message() string {
	return c.essence.Message()
}
