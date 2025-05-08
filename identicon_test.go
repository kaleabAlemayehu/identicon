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
