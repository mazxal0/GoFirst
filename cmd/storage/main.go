package main

import "fmt"
import "github.com/istrel/storage/internal/storage"

func main() {
	st := storage.NewStorage()

	fmt.Println("Hello World", st)
}
