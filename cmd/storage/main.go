package main

import (
	"fmt"

	"github.com/medvedevvu/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	fmt.Println("It works", st)
}
