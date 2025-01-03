package main

import (
	"dirmaker/cmd"
	"fmt"
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("banner.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))

	cmd.InitCobra()
}
