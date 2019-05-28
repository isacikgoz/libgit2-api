package git

import (
	"time"

	lib "gopkg.in/libgit2/git2go.v27"
)

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
func (r *Repository) Commit(message string, author ...*Signatute) error {
	repo := r.essence
	head, err := repo.Head()
	if err != nil {
		return err
	}
	defer head.Free()
	parent, err := repo.LookupCommit(head.Target())
	if err != nil {
		return err
	}
	defer parent.Free()
	index, err := repo.Index()
	if err != nil {
		return err
	}
	defer index.Free()
	if err := index.Write(); err != nil {
		return err
	}
	treeid, err := index.WriteTree()
	if err != nil {
		return err
	}
	tree, err := repo.LookupTree(treeid)
	if err != nil {
		return err
	}
	defer tree.Free()
	_, err = repo.CreateCommit("HEAD", author[0].toNewLibSignature(), author[0].toNewLibSignature(), message, tree, parent)
	if err != nil {
		return err
	}
	return nil
}
