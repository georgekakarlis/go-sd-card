package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
)

func NavigateFileSystem(startDir string) {
	currentDir := startDir
	for {
		entries, err := os.ReadDir(currentDir)
		if err != nil {
			fmt.Printf("Failed to read directory: %v\n", err)
			return
		}

		dirNames := []string{"[Exit]", "[Up One Level]"}
		for _, entry := range entries {
			dirNames = append(dirNames, entry.Name())
		}

		prompt := promptui.Select{
			Label: fmt.Sprintf("Current directory: %s", currentDir),
			Items: dirNames,
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		if result == "[Exit]" {
			return
		} else if result == "[Up One Level]" {
			currentDir = filepath.Dir(currentDir)
		} else {
			newDir := filepath.Join(currentDir, strings.Split(result, " (")[0]) // Extract only the name part
			selectedInfo, err := os.Stat(newDir)
			if err != nil {
				fmt.Printf("Failed to retrieve info: %v\n", err)
				continue
			}

			if selectedInfo.IsDir() {
				if confirmAction("Reorganize images in this directory by date? (Y/N)") {
					if err := ReorganizeImagesByDate(newDir); err != nil {
						fmt.Printf("Error reorganizing images: %v\n", err)
					} else {
						fmt.Println("Images successfully reorganized.")
						return
					}
				}
				currentDir = newDir
			} else {
				fmt.Printf("Selected file: %s\n", newDir)
				// Handle file selection here
			}
		}
	}
}

func confirmAction(actionDescription string) bool {
	prompt := promptui.Prompt{
		Label:     actionDescription,
		IsConfirm: true,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return false
	}

	return strings.ToLower(result) == "y"
}
