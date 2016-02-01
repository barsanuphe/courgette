// Package courgette contains everything necessary to manage a bunch of photos using their metadata.
// Useful if you organize your picture with a different folder for each month.
package courgette

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Collection represents a collection of Pictures
type Collection struct {
	Config
	Contents []SubDirectory
}

// Import from card.
func (c *Collection) Import(from string) (numImported int, err error) {
	incomingTarget := filepath.Join(c.Root, c.Incoming)
	// if it does not exist, create
	if _, err = os.Stat(incomingTarget); os.IsNotExist(err) {
		fmt.Println("Creating directory " + incomingTarget)
		os.MkdirAll(incomingTarget, 0700)
	}
	fmt.Println("Moving pictures to " + incomingTarget)
	// from -> c.Incoming
	// TODO walk from card, create Picture, copy
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
	subd, err := c.GetSubDir(subdir)
	if err != nil {
		return 0, err
	}
	numFiles, err = subd.Analyze(*c)
	return
}

// Refresh filenames in a given subdirectory.
func (c *Collection) Refresh(subdir string) (numRenamed int, err error) {
	subd, err := c.GetSubDir(subdir)
	if err != nil {
		return 0, err
	}
	// TODO loop over Pictures, BwPictures, RawFiles
	for _, pic := range subd.Jpgs {
		fmt.Println(pic.Filename)
		// TODO analyze, rename
	}
	return
}

// GetSubDir returns the SubDirectory from its name
func (c *Collection) GetSubDir(subdir string) (SubDirectory, error) {
	fullPath := filepath.Join(c.Root, subdir)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return SubDirectory{}, err
	}
	for _, subd := range c.Contents {
		if subd.Name == subdir {
			return subd, nil
		}
	}
	return SubDirectory{}, errors.New("Subdirectory " + subdir + " not found")
}
