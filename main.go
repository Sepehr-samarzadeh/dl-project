package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	content, err := http.Get("https://picsum.photos/200/300")

	if err != nil {
		panic(err)
	}
	defer content.Body.Close()

	file, err := os.Create("image.jpeg")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	_, err = io.Copy(file, content.Body)

	if err != nil {
		panic(err)
	}

}
