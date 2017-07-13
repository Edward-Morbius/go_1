package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println("First go.")

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	} else {

		for num, file := range files {
			fmt.Printf("%04x $%08X %52s\n", num, file.Size(), file.Name())
		}
	}
}
