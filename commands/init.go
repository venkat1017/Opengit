package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

func Init(args []string) {
	path := "."
	if len(args) > 0 {
		path = args[0]
		// Create the directory if it doesn't exist
		err := os.MkdirAll(path, 0755)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating directory %s: %v\n", path, err)
			os.Exit(1)
		}
	}

	// Create the .opengit directory instead of .git
	gitdir := filepath.Join(path, ".opengit")
	err := os.MkdirAll(gitdir, 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating .opengit directory: %v\n", err)
		os.Exit(1)
	}

	// Create the required subdirectories
	dirs := []string{"objects", "refs/heads", "refs/tags"}
	for _, dir := range dirs {
		err := os.MkdirAll(filepath.Join(gitdir, dir), 0755)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating %s: %v\n", dir, err)
			os.Exit(1)
		}
	}

	// Write the initial HEAD file
	err = os.WriteFile(filepath.Join(gitdir, "HEAD"), []byte("ref: refs/heads/main\n"), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing HEAD: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Initialized empty Opengit repository in %s\n", gitdir)
}
