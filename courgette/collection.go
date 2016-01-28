// Package courgette contains everything necessary to manage a bunch of photos using their metadata.
// Useful if you organize your picture with a different folder for each month.
package courgette

// Collection represents a collection of Pictures
type Collection struct {
	Config
	Contents []SubDirectory
}

// Import from card.
func (c *Collection) Import(from string) (numImported int, err error) {
	// from -> c.Incoming
	// copy then remove source
	return
}

// SortNew Pictures in the incoming directory.
func (c *Collection) SortNew() (numSorted int, err error) {
	//  c.Incoming -> p.NewPath
	return
}

// AnalyzeContents of a given subdirectory.
func (c *Collection) AnalyzeContents(subdir string) (numFiles int, err error) {
	// TODO find SubDirectory
	// TODO s.Analyze()
	return
}

// Refresh filenames in a given subdirectory.
func (c *Collection) Refresh(subdir string) (numRenamed int, err error) {
	return
}
