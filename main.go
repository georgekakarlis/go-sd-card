package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)



func main(){

	//mainDirectory := time.Now().Format("2006-01-02")


	// we get the current user logged in the device and his home directory, 99% is cached :/
    currentUser, err := user.Current()
    if err != nil {
       log.Fatalf(err.Error())
    }
    username := currentUser.Username
    userDir := currentUser.HomeDir
    fmt.Printf("Username is: %s\n", username)
    fmt.Printf("Home Directory is: %s\n", userDir)

	//the directory we want to create 
	sdCardNewFileDir := userDir + "/sdcardformat"
	// we try to create the directory
	err = os.Mkdir(sdCardNewFileDir , 0755)
	// we handle the error in case of one
	if err != nil {
		// Check if the error is because the directory already exists
		if os.IsExist(err) {
			fmt.Printf("Directory %s already exists\n", sdCardNewFileDir)
		} else {
			// If the error is of another kind, log it and exit
			log.Fatalf("Failed to create directory: %s", err)
		}
	} else {
		fmt.Printf("Directory %s created\n", sdCardNewFileDir)
	}


	
	// check os
	operatingSystem := runtime.GOOS
    switch operatingSystem {
    case "windows":
		runItOnWindows()
        fmt.Println("Windows operating system")
    case "darwin":
		runItOnMac()
        fmt.Println("OSX operating system")
    case "linux":
		runItOnLinux()
        fmt.Println("Linux operating system")
    default:
        fmt.Printf("Unsupported operating system: %s.\n", operatingSystem)
        os.Exit(1)
    }


}

func runItOnWindows(){
	//var mediaBasePathWindows string
}

func runItOnLinux(){

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

func runItOnMac(){
	
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
}


	//Scan for media mounted on the computer. Display a menu of storage devices. <== DO THIS ONE. DO A CLI

	//Select the device containing the media.
