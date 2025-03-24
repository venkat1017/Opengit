package commands

import (
	"Opengit/objects"
	"Opengit/repo"
	"fmt"
	"os"
	"path/filepath"
)

func HashObject(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: Opengit hash-object [-w] <file>")
		os.Exit(1)
	}

	write := false
	file := args[0]
	if args[0] == "-w" {
		if len(args) < 2 {
			fmt.Println("Usage: Opengit hash-object -w <file>")
			os.Exit(1)
		}
		write = true
		file = args[1]
	}

	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	blob := objects.NewBlob(data)
	hash := blob.Hash()
	fmt.Println(hash)

	if write {
		r, err := repo.NewRepository(".")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		objPath := r.GitFile("objects", hash[:2], hash[2:])
		err = os.MkdirAll(filepath.Dir(objPath), 0755)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating directory: %v\n", err)
			os.Exit(1)
		}
		err = os.WriteFile(objPath, blob.Serialize(), 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing object: %v\n", err)
			os.Exit(1)
		}
	}
}
