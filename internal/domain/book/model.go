package book

import "fmt"

type Book struct {
	UUID   string `json:"uuid,omitempty"`
	Name   string `json:"name,omitempty"`
	Year   int    `json:"year,omitempty"`
	Author string `json:"author,omitempty"`
	Busy   bool   `json:"busy,omitempty"`
	Owner  string `json:"owner,omitempty"`
}

func (b *Book) Take(owner string) error {
	if b.Busy {
		return fmt.Errorf("book is busy")
	}
	b.Owner = owner
	b.Busy = true
	return nil
}
