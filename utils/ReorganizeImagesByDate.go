package utils

import (
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/rwcarlsen/goexif/exif"
)

func ReorganizeImagesByDate(dirPath string) error {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue // Skip directories
		}

		filePath := filepath.Join(dirPath, entry.Name())
		takenDate := getPhotoTakenDate(filePath)

		if takenDate.IsZero() {
			continue // Skip files without a valid date
		}

		dateFolder := takenDate.Format("2006-01-02") // Format date as YYYY-MM-DD
		dateFolderPath := filepath.Join(dirPath, dateFolder)

		if err := os.MkdirAll(dateFolderPath, fs.ModePerm); err != nil {
			return err // Fail if unable to create a directory
		}

		newFilePath := filepath.Join(dateFolderPath, entry.Name())
		if err := os.Rename(filePath, newFilePath); err != nil {
			return err // Fail if unable to move the file
		}
	}

	return nil
}

func getPhotoTakenDate(filePath string) time.Time {
	f, err := os.Open(filePath)
	if err != nil {
		return time.Time{} // Return zero time on error
	}
	defer f.Close()

	x, err := exif.Decode(f)
	if err != nil {
		return time.Time{} // Return zero time on error
	}

	tm, err := x.DateTime()
	if err != nil {
		return time.Time{} // Return zero time if the date is not available
	}

	return tm
}
