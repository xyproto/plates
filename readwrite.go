package imagelib

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	ico "github.com/biessek/golang-ico"
	"github.com/chai2010/webp"
	bmp "github.com/jsummers/gobmp"
	"github.com/xyproto/xpm"
)

// Read tries to read the given image filename and return an image.Image
// The supported extensions are: .png, .jpg, .jpeg, .gif, .ico, .bmp and .webp
func Read(filename string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	switch strings.ToLower(filepath.Ext(filename)) {
	case ".png":
		return png.Decode(f)
	case ".jpg", ".jpeg":
		return jpeg.Decode(f)
	case ".gif":
		return gif.Decode(f)
	case ".ico":
		return ico.Decode(f)
	case ".bmp":
		return bmp.Decode(f)
	case ".webp":
		return webp.Decode(f)

	}
	return nil, errors.New("unrecognized file extension: " + filepath.Ext(filename))
}

// Write tries to write then given image.Image to a file.
// The supported extensions are: .png, .jpg, .jpeg. .gif, .ico, .bmp, .webp and .xpm
func Write(filename string, img image.Image) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	switch strings.ToLower(filepath.Ext(filename)) {
	case ".png":
		return png.Encode(f, img)
	case ".jpg", ".jpeg":
		return jpeg.Encode(f, img, nil)
	case ".gif":
		return gif.Encode(f, img, nil)
	case ".ico":
		return ico.Encode(f, img)
	case ".bmp":
		return bmp.Encode(f, img)
	case ".webp":
		return webp.Encode(f, img, nil)
	case ".xpm":
		return xpm.Encode(f, img)
	}
	return errors.New("unrecognized file extension: " + filepath.Ext(filename))
}
