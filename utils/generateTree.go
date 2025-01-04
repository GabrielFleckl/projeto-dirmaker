package utils

import (
	"fmt"
	"strings"
)

func GenerateTree(subDirs, files []string) {
	Tree := append(subDirs, files...)

	var filteredTree []string
	for _, item := range Tree {
		if strings.TrimSpace(item) != "" {
			filteredTree = append(filteredTree, item)
		}
	}

	for index, Branch := range filteredTree {
		if index < len(subDirs) {
			Branch += "/"
		}

		if index == len(filteredTree)-1 {
			fmt.Println("└──" + Branch)
		} else {
			fmt.Println("├──" + Branch)
		}
	}
}