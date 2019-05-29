package git

import (
	"strings"

	lib "gopkg.in/libgit2/git2go.v27"
)

// Branch is a wrapper of lib.Branch object
type Branch struct {
	refType RefType
	essence *lib.Branch
	owner   *Repository
	target  *Commit

	Name     string
	FullName string
	Hash     string
	isRemote bool
	Head     bool
	Ahead    int
	Behind   int
	Upstream *Branch
}

// Branches loads branches with the lib's branch iterator
// loads both remote and local branches
func (r *Repository) Branches() ([]*Branch, error) {
	branchIter, err := r.essence.NewBranchIterator(lib.BranchAll)
	if err != nil {
		return nil, err
	}
	defer branchIter.Free()
	buffer := make([]*Branch, 0)

	err = branchIter.ForEach(func(branch *lib.Branch, branchType lib.BranchType) error {
		b, err := unpackRawBranch(branch)
		if err != nil {
			return err
		}
		obj, err := r.essence.RevparseSingle(b.Hash)
		if err == nil && obj != nil {
			if commit, _ := obj.AsCommit(); commit != nil {
				b.target = unpackRawCommit(r, commit)
			}
		}
		buffer = append(buffer, b)
		return nil
	})

	return buffer, err
}

func unpackRawBranch(branch *lib.Branch) (*Branch, error) {
	name, err := branch.Name()
	if err != nil {
		return nil, err
	}
	fullname := branch.Reference.Name()

	rawOid := branch.Target()

	if rawOid == nil {
		ref, err := branch.Resolve()
		if err != nil {
			return nil, err
		}
		rawOid = ref.Target()
	}

	hash := rawOid.String()
	isRemote := branch.IsRemote()
	isHead, _ := branch.IsHead()

	var upstream *Branch
	if !isRemote {
		us, err := branch.Upstream()
		if err != nil || us == nil {
			// upstream not found
		} else {
			upstream = &Branch{
				Name:     strings.Replace(us.Name(), "refs/remotes/", "", 1),
				FullName: us.Name(),
				Hash:     us.Target().String(),
				isRemote: true,
				essence:  us.Branch(),
			}
		}
	}

	b := &Branch{
		Name:     name,
		refType:  RefTypeBranch,
		essence:  branch,
		FullName: fullname,
		Hash:     hash,
		isRemote: isRemote,
		Head:     isHead,
		Upstream: upstream,
	}
	return b, nil
}

// Type is the reference type of this ref
func (b *Branch) Type() RefType {
	return b.refType
}

// Target is the hash of targeted commit
func (b *Branch) Target() *Commit {
	return b.target
}
