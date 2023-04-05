package main

import "log"

func main() {
	if err := Chat(); err != nil {
		log.Fatal(err)
	}
}

// remove '//go:build mage' from the top of magefile.go
