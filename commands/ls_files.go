package commands

import (
	"Opengit/index"
	"Opengit/repo"
	"fmt"
	"os"
)

func LsFiles(args []string) {
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

	for path := range idx.Entries {
		fmt.Println(path)
	}
}
