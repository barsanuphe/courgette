package courgette

import "strings"

// Pictures is a list of Pictures.
type Pictures []Picture

func (p Pictures) Len() int {
	return len(p)
}

func (p Pictures) Less(i, j int) bool {
	return strings.ToLower(p[i].ID) < strings.ToLower(p[j].ID)
}

func (p Pictures) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// HasID checks if a list of Pictures contains one with a specific ID.
func (p *Pictures) HasID(id string) (hasID bool) {
	for _, pic := range *p {
		if pic.ID == id {
			hasID = true
			return
		}
	}
	return
}
