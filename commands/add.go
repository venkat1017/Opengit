package commands

import (
	"Opengit/index"
	"Opengit/objects"
	"Opengit/repo"
	"fmt"
	"os"
)

func Add(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: Opengit add <file>...")
		os.Exit(1)
	}

	r, err := repo.NewRepository(".")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	idx, err := index.ReadIndex(r)
	if err != nil {
		idx = &index.Index{Entries: make(map[string]*index.Entry)}
	}

	for _, file := range args {
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", file, err)
			continue
		}
		blob := objects.NewBlob(data)
		hash := blob.Hash()
		idx.Entries[file] = &index.Entry{Hash: hash}
	}

	err = index.WriteIndex(r, idx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing index: %v\n", err)
		os.Exit(1)
	}
}
