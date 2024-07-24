package imageloader

import (
	"image"
	"image/png"
	"io"
)

type PNGDecoder struct{}

func (d PNGDecoder) Decode(r io.Reader) (image.Image, error) {
	return png.Decode(r)
}
