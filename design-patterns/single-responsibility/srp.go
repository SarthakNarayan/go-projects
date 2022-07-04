package srp

import (
	"errors"
	"fmt"
)

// we are making entries private, if used outside of the package
// cannot be initialised outside the package
type Journal struct {
	entries []string
}

func (j *Journal) AddJournal(text string) {
	(*j).entries = append((*j).entries, text)
}

func (j *Journal) RemoveEntry(entry int) error {
	if len(j.entries) < entry {
		return errors.New("Out of bounds entry")
	}
	j.entries = append(j.entries[:entry], j.entries[entry+1:]...)
	return nil
}

func (j *Journal) ListEntries() {
	for i, journal := range j.entries {
		fmt.Printf("%d %v\n", i, journal)
	}
}
