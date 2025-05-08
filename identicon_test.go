package identicon

import (
	"testing"
)

func TestHash(t *testing.T) {
	testCase0 := "kaleab is gonna get hashed kaleab is gonna get hashed kaleab is gonna get hashed kaleab is gonna get hashed"
	testCase1 := "idfjdkfjadlkfjal kfajsflkajsflkajfla skfjdklfjakldfjalkdjfsdklfjalfa"
	render(testCase0)
	render(testCase1)
}

// func TestGenerateImage(t *testing.T) {
// 	testCase1 := "kaleab is gonna get hashed kaleab is gonna get hashed kaleab is gonna get hashed kaleab is gonna get hashed"
// 	generateImage(testCase1)
// }

// func render(data string) {
// 	h := sha1.New()
// 	h.Write([]byte(data))
// 	hVal := binary.LittleEndian.Uint64(h.Sum(nil)[0:8])
// 	colour := color.NRGBA{
// 		R: uint8(h),
// 		G: uint8(h >> 8),
// 		B: uint8(h >> 16),
// 		A: uint8(1),
// 	}
// 	h >>= 24
// 	img := image.NewPaletted(image.Rect(0, 0, 420, 420), color.Palette{color.NRGBA{0xf0, 0xf0, 0xf0, 0xff}, colour})
// }
