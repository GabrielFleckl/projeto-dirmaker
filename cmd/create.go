package cmd

import (
	"dirmaker/utils"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

var (
	createFiles   string
	createSubDirs string
)

func Create(cmd *cobra.Command, args []string) {
	directoryName := args[0]

	if len(args) == 0 {
		cmd.Usage()
		return
	}

	// Create Main Directory
	err := os.MkdirAll(directoryName, os.ModePerm)
	if err != nil {
		fmt.Printf("Error To Create Directory '%s': %v\n", directoryName, err)
		return
	}
	fmt.Println(directoryName + "/")

	// Create Sub Directories
	subDirs := strings.Split(createSubDirs, ",")
	for _, subDir := range subDirs {
		if subDir == "" {
			continue
		}

		subDirPath := filepath.Join(directoryName, strings.TrimSpace(subDir))
		err := os.MkdirAll(subDirPath, os.ModePerm)
		if err != nil {
			fmt.Printf("Error To Create Sub Directory `%s`: `%v\n`", subDirPath, err)
			return
		}
	}

	// Create Files
	files := strings.Split(createFiles, ",")

	for _, file := range files {
		if file == "" {
			continue
		}

		filePath := filepath.Join(directoryName, strings.TrimSpace(file))
		f, err := os.Create(filePath)
		if err != nil {
			fmt.Printf("Error to create file '%s': %v\n", filePath, err)
			continue
		}
		f.Close()
	}
	// Show Log of Created Branches
	utils.GenerateTree(subDirs, files)
}
