package main

import (
	"os"

	"github.com/kaleabAlemayehu/identicon"
)

func main() {
	file, _ := os.Create("./testfile.png")
	icon := identicon.New5X5()
	buf := icon.Render([]byte("test data for the file.go"))
	file.Write(buf)
}
