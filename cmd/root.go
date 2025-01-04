package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "dirmaker [main directory]",
	Short: "DirMaker is a very fast CLI directory and files structure creator.",
	Args:  cobra.ExactArgs(1),
	Run:   Create,
	PreRun: func(cmd *cobra.Command, args []string) {

		fmt.Println(" 📁 Structure created:")
		fmt.Println("")

	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("")
		fmt.Println("All Done! 👌🏻")
	},
}

// 	PostRun: func(cmd *cobra.Command, args []string) {
// 		fmt.Println("")
// 		fmt.Println("All Done!")
// 	},
// }

func InitCobra() {

	// add flags to rootCmd
	rootCmd.Flags().StringVarP(&createSubDirs, "subdirs", "s", "", "List of sub directories")

	rootCmd.Flags().StringVarP(&createFiles, "files", "f", "", "List of files")

	if err := rootCmd.Execute(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
