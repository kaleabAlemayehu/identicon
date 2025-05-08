package identicon

import (
	"os"
	"testing"
)

func TestHash(t *testing.T) {
	testCase0 := []byte("kaleab is gonna get hashed kaleab is gonna get hashed kaleab is gonna get hashed kaleab is gonna get hashed")
	testCase1 := []byte("idfjdkfjadlkfjal kfajsflkajsflkajfla skfjdklfjakldfjalkdjfsdklfjalfa")
	img, err := os.Create("./test0.png")
	if err != nil {
		t.Error(err.Error())
	}
	icon := New7X7()

	_, err = img.Write(icon.Render(testCase0))
	if err != nil {
		t.Error(err.Error())
	}
	img, err = os.Create("./test1.png")
	if err != nil {
		t.Error(err.Error())
	}
	_, err = img.Write(icon.Render(testCase1))
	if err != nil {
		t.Error(err.Error())
	}
}
