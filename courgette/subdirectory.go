package courgette

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"time"
)

// SubDirectory contains JPGs, B/W JPGs, and raw files.
type SubDirectory struct {
	Name     string
	Jpgs     Pictures
	BwJpgs   Pictures
	RawFiles Pictures
	Analyzed bool
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("-- [%s in %s]\n", name, elapsed)
}

// Analyze of a given subdirectory.
func (s *SubDirectory) Analyze(c Collection) (numFiles int, err error) {
	// TODO use interface Config instead of Collection
	defer timeTrack(time.Now(), "Scanned")

	filepath.Walk(s.Name, func(path string, f os.FileInfo, err error) (outErr error) {
		if f.Mode().IsRegular() {
			fileName := filepath.Base(path)
			found := reg.FindStringSubmatch(fileName)
			if len(found) == 6 {
				var isBW bool
				number, _ := strconv.Atoi(found[2])
				if found[3] != "" {
					isBW = true
				}
				version, _ := strconv.Atoi(found[4])
				extension := found[5]
				id := found[2] + found[1] // number + prefix

				p, err := NewPicture(fileName, isBW, number, version, id, extension)
				if err != nil {
					return
				}
				switch extension {
				case ".cr2", ".arw":
					s.RawFiles = append(s.RawFiles, *p)
				case ".jpg":
					if isBW {
						s.BwJpgs = append(s.BwJpgs, *p)
					} else {
						s.Jpgs = append(s.Jpgs, *p)
					}
				}
				numFiles++
			} else {
				fmt.Println("Error with ", fileName)
			}
		}
		return
	})
	fmt.Printf("Found %d files.\n", numFiles)
	s.Analyzed = true
	return
}

// FindOrphans in a given subdirectory
func (s *SubDirectory) FindOrphans(c Collection) (orphans Pictures, err error) {
	defer timeTrack(time.Now(), "Analyzed")

	// Analyze if necessary
	if !s.Analyzed {
		_, err = s.Analyze(c)
		if err != nil {
			return Pictures{}, err
		}
	}

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
