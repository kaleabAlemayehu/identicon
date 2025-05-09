package main

import (
	"github.com/kaleabAlemayehu/identicon"
	"log"
	"net/http"
	"os"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {

	args := strings.Split(r.URL.Path, "/")
	args = args[1:]

	if len(args) != 1 {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	item := args[0]

	if !strings.HasSuffix(item, ".png") {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	item = strings.TrimSuffix(item, ".png")

	icon := identicon.New7X7()

	log.Printf("creating identicon for '%s'\n", item)

	data := []byte(item)
	pngdata := icon.Render(data)

	w.Header().Set("Content-Type", "image/png")
	w.Write(pngdata)

	return
}

func main() {
	port := ":8080"
	if p := os.Getenv("PORT"); p != "" {
		port = ":" + p
	}
	log.Println("Listening on port", port)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(port, nil))
}
