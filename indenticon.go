package identicon

import (
	"crypto/sha1"
	"encoding/binary"
	"fmt"

	"image"
	"image/color"
	"image/png"

	// "io"
	"log"
	"math/rand"
	"os"
	"time"
)

func hashit(data string) []byte {
	h := sha1.New()
	h.Write([]byte(data))
	hVal := binary.BigEndian.Uint64(h.Sum(nil)[0:8])
	fmt.Println(hVal)
	return h.Sum(nil)
}
func randomColor() color.RGBA {
	rand := rand.New(rand.NewSource(time.Now().Unix()))
	return color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}
}

func generateImage(data string) {
	img := image.NewPaletted(image.Rect(0, 0, 360, 360), color.Palette{color.NRGBA{0x0f, 0x0f, 0x0f, 0xff}})
	for x := 0; x < 160; x++ {
		for y := 0; y < 160; y++ {
			img.Set(x, y, randomColor())
		}
	}
	i, err := os.Create(fmt.Sprintf("./test%s.png", string(data[10])))
	if err != nil {
		log.Println(err.Error())
	}
	png.Encode(i, img)
}

func render(data string) {
	h := sha1.New()
	h.Write([]byte(data))
	hVal := binary.LittleEndian.Uint64(h.Sum(nil)[0:8])
	colour := color.NRGBA{
		R: uint8(hVal),
		G: uint8(hVal >> 8),
		B: uint8(hVal >> 16),
		A: uint8(1),
	}

	hVal >>= 24

	const xborder = 35
	const yborder = 35
	const maxX = 420
	const maxY = 420

	sqx := 0
	sqy := 0

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

	i, err := os.Create(fmt.Sprintf("./test%s.png", string(data[10])))
	if err != nil {
		log.Println(err.Error())
	}
	png.Encode(i, img)
}
