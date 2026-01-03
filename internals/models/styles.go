package models

import (
	"log"
	"os"

	"golang.org/x/term"
)

func GetTermSize() (int, int) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Fatalf("Something went wrong while finding size: %s", err)
	}
	return width, height
}
