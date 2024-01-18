package gosdcard

import (
	"fmt"
	"gosdcard/pkg"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gosdcard",
	Short: "GoSDCard is a tool for managing Disks/File System",
	Long: `A Fast and Flexible Disk/File management tool built with 
           love by Georgios in Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		pkg.MainLogic()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
