package pkg

import (
	"fmt"
	"gosdcard/utils"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
)

func MainLogic() {
	showLoadingScreen()
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	username := currentUser.Username
	userDir := currentUser.HomeDir
	fmt.Printf("Username is: %s\n", username)
	fmt.Printf("Home Directory is: %s\n", userDir)

	sdCardNewFileDir := filepath.Join(userDir, "sdcardformat")
	err = os.Mkdir(sdCardNewFileDir, 0755)
	if err != nil {
		if os.IsExist(err) {
			fmt.Printf("Directory %s already exists\n", sdCardNewFileDir)
		} else {
			log.Fatalf("Failed to create directory: %s", err)
		}
	} else {
		fmt.Printf("Directory %s created\n", sdCardNewFileDir)
	}

	switch runtime.GOOS {
	case "windows":
		fmt.Println("Windows operating system")
		runItOnWindows()
	case "darwin":
		fmt.Println("OSX operating system")
		runItOnMac()
	case "linux":
		fmt.Println("Linux operating system")
		runItOnLinux()
	default:
		fmt.Printf("Unsupported operating system: %s.\n", runtime.GOOS)
		os.Exit(1)
	}
}

func runItOnWindows() {
	// Windows-specific code
}

func runItOnLinux() {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	mediaBasePathLinux := filepath.Join("/media", currentUser.Username)
	mediaDirs, err := os.ReadDir(mediaBasePathLinux)
	if err != nil {
		log.Fatalf("Failed to read the media directory: %v", err)
	}

	if len(mediaDirs) == 0 {
		fmt.Println("No storage media found.")
	} else {
		for i, dir := range mediaDirs {
			fmt.Printf("%d. %s\n", i+1, filepath.Join(mediaBasePathLinux, dir.Name()))
		}
	}
}

func runItOnMac() {
	mediaBasePathMac := "/Volumes"
	mediaDirs, err := os.ReadDir(mediaBasePathMac)
	if err != nil {
		log.Fatalf("Failed to read the media directory: %v", err)
	}

	if len(mediaDirs) == 0 {
		fmt.Println("No storage media found.")
	} else {
		for i, dir := range mediaDirs {
			fmt.Printf("%d. %s\n", i+1, filepath.Join(mediaBasePathMac, dir.Name()))
		}
	}
	driveNames := []string{}
	for _, dir := range mediaDirs {
		driveNames = append(driveNames, filepath.Join(mediaBasePathMac, dir.Name()))
	}

	// Let the user choose a drive
	selectedDrive := chooseDrive(driveNames)
	utils.NavigateFileSystem(selectedDrive)
}

func chooseDrive(drives []string) string {
	prompt := promptui.Select{
		Label: "Select Drive",
		Items: drives,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}

func showLoadingScreen() {
	for i := 0; i < 10; i++ {
		fmt.Printf("\rLoading %s", strings.Repeat(".", i%4))
		time.Sleep(250 * time.Millisecond)
	}
	fmt.Println("\rLoading complete!      ") // Clear the loading text
}
