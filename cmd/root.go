package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "dirmaker",
	Short: "DirMaker is a very fast CLI directory and files structure creator.",
}

var createCmd = &cobra.Command{
	Use:   "create [main directory]",
	Short: "Create directories and files with args",
	Args:  cobra.ExactArgs(1),
	Run:   Create,
	PreRun: func(cmd *cobra.Command, args []string) {

		fmt.Println("Creating...")
		fmt.Println("")

	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("")
		fmt.Println("All Done!")
	},
}

func InitCobra() {

	// add flags to cmd "create"
	createCmd.Flags().StringVarP(&createSubDirs, "subdirs", "s", "", "List of sub directories")

	createCmd.Flags().StringVarP(&createFiles, "files", "f", "", "List of files")

	// add create to dirmaker
	rootCmd.AddCommand(createCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
