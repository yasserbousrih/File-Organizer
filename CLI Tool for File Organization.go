package main

import (
	"fmt"           // For printing messages
	"os"            // For file and directory operations
	"path/filepath" // For handling file paths
)

func main() {
	sourceDir := "./" // Specify the directory to organize

	// Walk through all files in the directory
	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // Handle any errors during file traversal
		}
		// If it's a file (not a directory)
		if !info.IsDir() {
			ext := filepath.Ext(info.Name()) // Get the file's extension
			// Create a target directory based on the extension
			targetDir := filepath.Join(sourceDir, ext[1:]) // Remove the dot from the extension
			// Check if the target directory exists; if not, create it
			if _, err := os.Stat(targetDir); os.IsNotExist(err) {
				os.Mkdir(targetDir, 0755) // Create the directory with appropriate permissions
			}
			// Move the file to the target directory
			newPath := filepath.Join(targetDir, info.Name())
			os.Rename(path, newPath)
		}
		return nil
	})

	// Check for errors during the process
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Files organized successfully!")
	}
}
