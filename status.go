package git

import (
	lib "gopkg.in/libgit2/git2go.v27"
)

// State is the current state of the repository
type State int

// The different states for a repo
const (
	StateUnknown State = iota
	StateNone
	StateMerge
	StateRevert
	StateCherrypick
	StateBisect
	StateRebase
	StateRebaseInteractive
	StateRebaseMerge
	StateApplyMailbox
	StateApplyMailboxOrRebase
)

// IndexType describes the different stages a status entry can be in
type IndexType int

// The different status stages
const (
	IndexTypeStaged IndexType = iota
	IndexTypeUnstaged
	IndexTypeUntracked
	IndexTypeConflicted
)

// StatusEntryType describes the type of change a status entry has undergone
type StatusEntryType int

// The set of supported StatusEntryTypes
const (
	StatusEntryTypeNew StatusEntryType = iota
	StatusEntryTypeModified
	StatusEntryTypeDeleted
	StatusEntryTypeRenamed
	StatusEntryTypeUntracked
	StatusEntryTypeTypeChange
	StatusEntryTypeConflicted
)

var indexTypeMap = map[lib.Status]IndexType{
	lib.StatusIndexNew | lib.StatusIndexModified | lib.StatusIndexDeleted | lib.StatusIndexRenamed | lib.StatusIndexTypeChange: IndexTypeStaged,
	lib.StatusWtModified | lib.StatusWtDeleted | lib.StatusWtTypeChange | lib.StatusWtRenamed:                                  IndexTypeUnstaged,
	lib.StatusWtNew:      IndexTypeUntracked,
	lib.StatusConflicted: IndexTypeConflicted,
}

var statusEntryTypeMap = map[lib.Status]StatusEntryType{
	lib.StatusIndexNew:        StatusEntryTypeNew,
	lib.StatusIndexModified:   StatusEntryTypeModified,
	lib.StatusWtModified:      StatusEntryTypeModified,
	lib.StatusIndexDeleted:    StatusEntryTypeDeleted,
	lib.StatusWtDeleted:       StatusEntryTypeDeleted,
	lib.StatusIndexRenamed:    StatusEntryTypeRenamed,
	lib.StatusWtRenamed:       StatusEntryTypeRenamed,
	lib.StatusIndexTypeChange: StatusEntryTypeTypeChange,
	lib.StatusWtTypeChange:    StatusEntryTypeTypeChange,
	lib.StatusWtNew:           StatusEntryTypeUntracked,
	lib.StatusConflicted:      StatusEntryTypeConflicted,
}

// StatusEntry contains data for a single status entry
type StatusEntry struct {
	index           IndexType
	statusEntryType StatusEntryType
	diffDelta       *DiffDelta
}

// Status contains all git status data
type Status struct {
	State    State
	Entities []*StatusEntry
}

// Diff is the wrapper for a diff content acquired from repo
type Diff struct {
	deltas []*DiffDelta
	stats  []string
	patchs []string
}

// Deltas returns the actual changes with file info
func (d *Diff) Deltas() []*DiffDelta {
	return d.deltas
}

// DiffDelta holds delta status, file changes and the actual patchs
type DiffDelta struct {
	Status  int
	OldFile *DiffFile
	NewFile *DiffFile
	Patch   string
}

// DiffFile the file that has been changed
type DiffFile struct {
	Path string
	Hash string
}

func (r *Repository) LoadStatus() (*Status, error) {
	statusOptions := &lib.StatusOptions{
		Show:  lib.StatusShowIndexAndWorkdir,
		Flags: lib.StatusOptIncludeUntracked,
	}
	statusList, err := r.essence.StatusList(statusOptions)
	if err != nil {
		return nil, err
	}
	defer statusList.Free()

	count, err := statusList.EntryCount()
	if err != nil {
		return nil, err
	}
	entities := make([]*StatusEntry, 0)
	s := &Status{
		State:    State(r.essence.State()),
		Entities: entities,
	}
	for i := 0; i < count; i++ {
		statusEntry, err := statusList.ByIndex(i)
		if err != nil {
			return nil, err
		}
		if statusEntry.Status <= 0 {
			continue
		}
		s.addToStatus(statusEntry)
	}
	return s, nil
}

func (s *Status) addToStatus(raw lib.StatusEntry) {
	for rawStatus, indexType := range indexTypeMap {
		set := raw.Status & rawStatus

		if set > 0 {
			var dd lib.DiffDelta
			if indexType == IndexTypeStaged {
				dd = raw.HeadToIndex
			} else {
				dd = raw.IndexToWorkdir
			}
			d := &DiffDelta{
				Status: int(dd.Status),
				NewFile: &DiffFile{
					Path: dd.NewFile.Path,
				},
				OldFile: &DiffFile{
					Path: dd.OldFile.Path,
				},
			}
			e := &StatusEntry{
				index:           indexType,
				statusEntryType: statusEntryTypeMap[set],
				diffDelta:       d,
			}
			s.Entities = append(s.Entities, e)
		}
	}
}

// Indexed true if entry added to index
func (e *StatusEntry) String() string {
	return e.diffDelta.OldFile.Path
}

// Indexed true if entry added to index
func (e *StatusEntry) Indexed() bool {
	return e.index == IndexTypeStaged
}

// StatusEntryString returns entry status in pretty format
func (e *StatusEntry) StatusEntryString() string {
	switch e.statusEntryType {
	case StatusEntryTypeNew:
		return "Added"
	case StatusEntryTypeDeleted:
		return "Deleted"
	case StatusEntryTypeModified:
		return "Modified"
	case StatusEntryTypeRenamed:
		return "Renamed"
	case StatusEntryTypeUntracked:
		return "Untracked"
	case StatusEntryTypeTypeChange:
		return "Type change"
	case StatusEntryTypeConflicted:
		return "Conflicted"
	default:
		return "Unknown"
	}
}

// AddToIndex is the wrapper of "git add /path/to/file" command
func (r *Repository) AddToIndex(e *StatusEntry) error {
	index, err := r.essence.Index()
	if err != nil {
		return err
	}
	if err := index.AddByPath(e.diffDelta.OldFile.Path); err != nil {
		return err
	}
	defer index.Free()
	return index.Write()
}

// RemoveFromIndex is the wrapper of "git reset path/to/file" command
func (r *Repository) RemoveFromIndex(e *StatusEntry) error {
	if !e.Indexed() {
		return ErrEntryNotIndexed
	}
	index, err := r.essence.Index()
	if err != nil {
		return err
	}
	if err := index.RemoveByPath(e.diffDelta.OldFile.Path); err != nil {
		return err
	}
	defer index.Free()
	return index.Write()
}
