package main

import (
	"fmt"
	"log"
)
import "github.com/istrel/storage/internal/storage"

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("Hello World!"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("It was uploaded", file)
}
