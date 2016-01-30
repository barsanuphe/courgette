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

	"github.com/barsanuphe/goexiftool"
)

// matches the type of filename I like best.
var reg = regexp.MustCompile(`^(.*?)_(\d{4,5})(-bw\d*)?(-\d*)?(\.jpg|\.cr2|\.mov|\.arw|\.mp4)$`)

// Picture can manipulate a picture file.
type Picture struct {
	goexiftool.MediaFile
	Hash    string
	NewPath string
	IsBW    bool
	Number  int
	Version int
	ID      string
}

// Rename a Picture from metadata.
func (p *Picture) Rename() (err error) {
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
