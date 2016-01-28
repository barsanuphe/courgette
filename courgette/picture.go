package courgette

import (
	"regexp"

	"github.com/barsanuphe/goexiftool"
)

// matches the type of filename I like best.
var reg = regexp.MustCompile(`^(.*?)_(\d{4,5})(-bw\d*)?(-\d*)?(\.jpg|\.cr2|\.mov|\.arw|\.mp4)$`)

// Picture can manipulate a picture file.
type Picture struct {
	goexiftool.MediaFile
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
func (p *Picture) IsNew(c Config) (isNew bool, err error) {
	return
}

// Rotate losslessly the Picture.
func (p *Picture) Rotate() (err error) {
	return
}
