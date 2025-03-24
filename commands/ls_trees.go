package commands

import (
	"Opengit/objects"
	"Opengit/repo"
	"bytes"
	"fmt"
	"os"
)

func LsTree(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: mygit ls-tree <tree-ish>")
		os.Exit(1)
	}

	r, err := repo.NewRepository(".")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	hash := args[0]
	objPath := r.GitFile("objects", hash[:2], hash[2:])
	data, err := os.ReadFile(objPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading tree %s: %v\n", hash, err)
		os.Exit(1)
	}

	if !bytes.HasPrefix(data, []byte("tree")) {
		fmt.Fprintf(os.Stderr, "Not a tree object: %s\n", hash)
		os.Exit(1)
	}

	content := data[bytes.IndexByte(data, '\x00')+1:]
	tree := objects.ParseTree(content)
	for _, entry := range tree.Entries {
		fmt.Printf("%s %s %s\t%s\n", entry.Mode, "blob", entry.Hash, entry.Name)
	}
}
