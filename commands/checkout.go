package commands

import (
	"Opengit/objects"
	"Opengit/refs"
	"Opengit/repo"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

func Checkout(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: Opengit checkout <commit-ish>")
		os.Exit(1)
	}

	r, err := repo.NewRepository(".")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	refs := refs.NewRefs(r)
	target := args[0]
	hash := target // Assume it’s a hash; resolve if it’s a ref
	if refHash, err := refs.ReadRef("refs/heads/" + target); err == nil && refHash != "" {
		hash = refHash
	}

	// Load commit and its tree
	commitData, err := r.ReadObject(hash)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading commit %s: %v\n", hash, err)
		os.Exit(1)
	}
	commit := objects.ParseCommit(commitData[bytes.IndexByte(commitData, '\x00')+1:])
	treeData, err := r.ReadObject(commit.Tree)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading tree %s: %v\n", commit.Tree, err)
		os.Exit(1)
	}
	tree := objects.ParseTree(treeData[bytes.IndexByte(treeData, '\x00')+1:])

	// Update working tree
	for _, entry := range tree.Entries {
		blobData, err := r.ReadObject(entry.Hash)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading blob %s: %v\n", entry.Hash, err)
			continue
		}
		content := blobData[bytes.IndexByte(blobData, '\x00')+1:]
		path := filepath.Join(r.Worktree, entry.Name)
		os.MkdirAll(filepath.Dir(path), 0755)
		os.WriteFile(path, content, 0644)
	}

	// Update HEAD
	refs.WriteHEAD(hash)
	fmt.Printf("Checked out %s\n", hash)
}
