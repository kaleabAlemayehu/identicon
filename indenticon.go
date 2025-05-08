package identicon

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"hash"

	"image"
	"image/color"
	"image/png"
)

type Identicon struct {
	sqSize int
	rows   int
	cols   int
	h      hash.Hash
}

func New5X5() *Identicon {
	return &Identicon{
		sqSize: 70,
		rows:   5,
		cols:   5,
		h:      sha1.New(),
	}
}

func New7X7() *Identicon {
	return &Identicon{
		sqSize: 50,
		rows:   7,
		cols:   7,
		h:      sha1.New(),
	}
}

func (i *Identicon) Render(data []byte) []byte {
	i.h.Write(data)
	hVal := binary.BigEndian.Uint64(i.h.Sum(nil)[0:8])
	colour := color.NRGBA{
		R: uint8(hVal),
		G: uint8(hVal >> 8),
		B: uint8(hVal >> 16),
		A: 0xff,
	}

	hVal >>= 24

	sqx := 0
	sqy := 0

	const xborder = 35
	const yborder = 35
	const maxX = 420
	const maxY = 420
	sqSize := 50
	rows := 7
	cols := 7

	pixels := make([]byte, sqSize)
	for i := 0; i < sqSize; i++ {
		pixels[i] = 1
	}
	img := image.NewPaletted(image.Rect(0, 0, 420, 420), color.Palette{color.NRGBA{0xf0, 0xf0, 0xf0, 0xff}, colour})

	for i := 0; i < rows*(cols+1)/2; i++ {
		if hVal&1 == 1 {
			for i := 0; i < sqSize; i++ {
				x := xborder + sqx*sqSize
				y := yborder + sqy*sqSize + i
				offs := img.PixOffset(x, y)
				copy(img.Pix[offs:], pixels)

				x = xborder + (cols-1-sqx)*sqSize
				offs = img.PixOffset(x, y)
				copy(img.Pix[offs:], pixels)
			}
		}
		hVal >>= 1
		sqy++
		if sqy == rows {
			sqy = 0
			sqx++
		}
	}

	var buf bytes.Buffer
	png.Encode(&buf, img)
	return buf.Bytes()
}
