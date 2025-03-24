package commands

import (
	"Opengit/index"
	"Opengit/repo"
	"fmt"
	"os"
)

func Rm(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: mygit rm <file>...")
		os.Exit(1)
	}

	r, err := repo.NewRepository(".")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	idx, err := index.ReadIndex(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading index: %v\n", err)
		os.Exit(1)
	}

	for _, file := range args {
		if _, exists := idx.Entries[file]; !exists {
			fmt.Fprintf(os.Stderr, "Error: %s not in index\n", file)
			continue
		}
		delete(idx.Entries, file)
	}

	err = index.WriteIndex(r, idx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing index: %v\n", err)
		os.Exit(1)
	}
}
