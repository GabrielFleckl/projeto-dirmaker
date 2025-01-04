package utils

import (
	"fmt"
	"log"
	"os"
)

func ShowBanner() {
	banner, err := os.ReadFile("./misc/banner.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(banner))
}
