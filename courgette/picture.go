package courgette

import (
	"crypto/sha512"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"

	"fmt"

	"github.com/barsanuphe/goexiftool"
)

// matches the type of filename I like best.
var reg = regexp.MustCompile(`^(.*?)_(\d{4,5})(-bw\d*)?(-\d*)?(\.jpg|\.cr2|\.mov|\.arw|\.mp4)$`)

// Picture can manipulate a picture file.
type Picture struct {
	goexiftool.MediaFile
	// TODO instead of Parent, subdir ie filepath.Rel(root, parent(filename))
	Parent        string
	NewFilename   string
	Hash          string
	IsBW          bool
	Number        int
	Version       int
	FormattedDate string
	ID            string
	Camera        string
	Lens          string
	Extension     string
	Analyzed      bool
}

// NewPicture initializes a Picture and parses its metadata with exiftool.
func NewPicture(filename string, isBW bool, number int, version int, id string, extension string) (p *Picture, err error) {
	if _, err = os.Stat(filename); os.IsNotExist(err) {
		return nil, err
	}
	return &Picture{
		goexiftool.MediaFile{Filename: filename, Info: make(map[string]string)},
		filepath.Dir(filename),
		"", "", isBW, number, version, "", id, "", "", extension, false,
	}, err
}

// Analyze a Picture to get its metadata and filename information.
func (p *Picture) Analyze(c Collection) (err error) {
	// TODO use interface Config instead of Collection
	err = p.AnalyzeMetadata()
	if err != nil {
		return
	}
	// get camera shortname, according to configuration
	camera, err := p.GetCamera()
	if err != nil {
		p.Camera = "UnknownCamera"
	} else {
		cameraShortName, ok := c.Cameras[camera]
		if !ok {
			p.Camera = camera
		} else {
			p.Camera = cameraShortName
		}
	}
	// get lens shortname, according to configuration
	lens, err := p.GetLens()
	if err != nil {
		p.Lens = "UnknownLens"
	} else {
		lensShortName, ok := c.Lenses[lens]
		if !ok {
			p.Lens = lens
		} else {
			p.Lens = lensShortName
		}
	}
	date, err := p.GetDate()
	if err != nil {
		p.FormattedDate = "UnknownDate"
	} else {
		// TODO: use format from config
		p.FormattedDate = date.Format("2006-01-02-15h04m05s")
	}
	err = p.ComputeHash()
	if err != nil {
		return
	}
	p.Analyzed = true
	return
}

// Rename a Picture from metadata.
func (p *Picture) Rename(c Collection) (wasRenamed bool, err error) {
	// TODO use interface Config instead of Collection
	if !p.Analyzed {
		err = p.Analyze(c)
		if err != nil {
			return
		}
	}
	// TODO : what if Analyze shows that we're in the wrong subdir?
	// TODO: which will happen if the file is in INCOMING
	// Mon Jan 2 15:04:05 -0700 MST 2006
	// TODO: new SubdirName

	// new Filename
	p.NewFilename = fmt.Sprintf("%s_%s.%s_%d", p.FormattedDate, p.Camera, p.Lens, p.Number)
	if p.IsBW {
		p.NewFilename += "-bw"
	}
	if p.Version != 0 {
		p.NewFilename += fmt.Sprintf("-%d", p.Version)
	}
	p.NewFilename += p.Extension
	// TODO: filepath.Join(Root, newSubdirName, p.NewFilename)
	newPath := filepath.Join(p.Parent, p.NewFilename)
	if p.Filename != newPath {
		fmt.Printf("Renaming: %s -> %s.\n", p.Filename, p.NewFilename)
		// rename
		os.Rename(p.Filename, newPath)
		wasRenamed = true

	}
	return
}

// ConvertToBW a color Picture.
func (p *Picture) ConvertToBW() (err error) {
	return
}

// IsNew is true if it is.
func (p *Picture) IsNew(c Config) (isNew bool) {
	return strings.Contains(p.Filename, filepath.Join(c.Root, c.Incoming))
}

// Rotate losslessly the Picture.
func (p *Picture) Rotate() (err error) {
	return
}

// ComputeHash calculates the hash of the Picture file.
func (p *Picture) ComputeHash() (err error) {
	var result []byte
	file, err := os.Open(p.Filename)
	if err != nil {
		return err
	}
	defer file.Close()

	pictureHash := sha512.New()
	if _, err := io.Copy(pictureHash, file); err != nil {
		return err
	}
	p.Hash = hex.EncodeToString(pictureHash.Sum(result))
	return
}

// Diff compares two Pictures.
func (p *Picture) Diff(otherP Picture) (isSame bool, diffText string, err error) {
	if p.Hash == otherP.Hash && reflect.DeepEqual(p.Info, otherP.Info) {
		return true, "", nil
	}
	// TODO diff
	return
}
