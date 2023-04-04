package main

import "log"

func main() {
	if err := Chat(); err != nil {
		log.Fatal(err)
	}
}

// cp magefile.go debug_magefile.go
// remove '//go:build mage' from the top of the file
