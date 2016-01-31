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
func NewPicture(filename string) (p *Picture, err error) {
	if _, err = os.Stat(filename); os.IsNotExist(err) {
		return nil, err
	}
	p = &Picture{
		goexiftool.MediaFile{Filename: filename, Info: make(map[string]string)},
		filepath.Dir(filename),
		"", "", false, 0, 0, "", "", "", "", "", false,
	}
	return
}

// Analyze a Picture to get its metadata and filename information.
func (p *Picture) Analyze(c Config) (err error) {
	err = p.AnalyzeMetadata()
	if err != nil {
		return
	}
	err = p.ComputeHash()
	if err != nil {
		return
	}
	// TODO: parse filename too + generate FormattedDate, ID, Camera, Lens with Config
	p.Analyzed = true
	return
}

// Rename a Picture from metadata.
func (p *Picture) Rename(c Config) (err error) {
	if !p.Analyzed {
		err = p.Analyze(c)
		if err != nil {
			return
		}
	}
	p.NewFilename = fmt.Sprintf("%s_%s.%s_%d", p.FormattedDate, p.Camera, p.Lens, p.Number)
	if p.IsBW {
		p.NewFilename += "-bw"
	}
	if p.Version != 0 {
		p.NewFilename += fmt.Sprintf("-%d", p.Version)
	}
	p.NewFilename += p.Extension
	fmt.Printf("Renaming: %s to %s.\n", p.Filename, p.NewFilename)

	// TODO set NewPath and os.Rename to it
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
