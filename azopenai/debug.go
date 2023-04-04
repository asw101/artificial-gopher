package main

import "log"

func main() {
	if err := Chat(); err != nil {
		log.Fatal(err)
	}
}
