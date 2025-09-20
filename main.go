package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// var i used count total files
// var c used to count total files changed
var i, c int = 0, 0

// replaceStringInFiles replace a old string to the new in files inside path
func replaceStringInFiles(root string, oldStr, newStr string) error {
	// Walk run the tree of files and directories
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Read content file
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// Convert content to string to make the search
		fileContent := string(content)

		// Verify if the old string exists in the file
		if strings.Contains(fileContent, oldStr) {
			fmt.Printf("%s \033[32m✓\033[0m\n", path)

			// Make replace
			newContent := strings.ReplaceAll(fileContent, oldStr, newStr)

			// Write the new content to the file, rewrite the old content (be carefull)
			if err := os.WriteFile(path, []byte(newContent), info.Mode()); err != nil {
				return err
			}
			// count total files changed
			c++
		}
		// count total files scanned
		i++
		// show file name for each file read <optional>
		// fmt.Println("File name:", path)
		return nil
	})

	fmt.Println("Total files scanned:", i)
	fmt.Println("Total files changed:", c)
	return nil
}

func main() {
	program_name := filepath.Base(os.Args[0])
	// Test arguments
	if len(os.Args) < 4 {
		fmt.Printf("Usage: %s <path> <old_string> <new_string>\n", program_name)
		os.Exit(1)
	}

	// Arguments
	directory := os.Args[1]
	oldString := os.Args[2]
	newString := os.Args[3]

	// Confirmation
	var confirm string
	fmt.Printf("\033[38;5;208mATTENTION:\033[0m Do you really want to replace '%s' to '%s' in %s? [y/N] ", oldString, newString, directory)
	fmt.Scanln(&confirm)
	if strings.ToLower(confirm) != "y" {
		os.Exit(1)
	}

	// Execute replace
	if err := replaceStringInFiles(directory, oldString, newString); err != nil {
		log.Fatalf("Error to replace files: %v", err)
	}

	fmt.Println("Proccess finished")
}
