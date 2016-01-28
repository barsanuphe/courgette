package courgette

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"
)

// SubDirectory contains JPGs, B/W JPGs, and raw files.
type SubDirectory struct {
	Name     string
	Jpgs     Pictures
	BwJpgs   Pictures
	RawFiles Pictures
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("-- [%s in %s]\n", name, elapsed)
}

// Analyze of a given subdirectory.
func (s *SubDirectory) Analyze() (numFiles int, err error) {
	defer timeTrack(time.Now(), "Scanned")
	count := 0

	filepath.Walk(s.Name, func(path string, f os.FileInfo, err error) (outErr error) {
		if f.Mode().IsRegular() {
			fileName := filepath.Base(path)

			found := reg.FindStringSubmatch(fileName)
			if len(found) == 6 {
				id := found[2] + found[1] // number + prefix

				if found[5] == ".cr2" || found[5] == ".arw" {
					// TODO create Picture with ID, etc, append to RawFiles
					p := Picture{ID: id}
					s.RawFiles = append(s.RawFiles, p)
				} else if found[5] == ".jpg" {
					// TODO create Picture with ID
					// TODO if BW, append to BwJpgs
					// TODO else, append to Jpgs
				}
				count++
			} else {
				fmt.Println("Error with ", fileName)
			}
		}
		return
	})
	fmt.Printf("Found %d files.\n", count)
	return
}

// FindOrphans in a given subdirectory
func (s *SubDirectory) FindOrphans() (orphans Pictures, err error) {
	defer timeTrack(time.Now(), "Analyzed")

	// TODO if not done: Analyze(subdir)

	for _, raw := range s.RawFiles {
		// search in Jpgs
		found := s.Jpgs.HasID(raw.ID)
		if !found {
			// if not found, search in BwJpgs
			found = s.BwJpgs.HasID(raw.ID)
		}

		if !found {
			orphans = append(orphans, raw)
		}
	}
	sort.Sort(orphans)
	return
}
