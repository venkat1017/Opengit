package commands

import (
	"Opengit/index"
	"Opengit/objects"
	"Opengit/refs"
	"Opengit/repo"
	"fmt"
	"os"
	"time"
)

func Commit(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: Opengit commit -m <message>")
		os.Exit(1)
	}

	if args[0] != "-m" {
		fmt.Println("Error: missing commit message")
		os.Exit(1)
	}

	message := args[1]
	r, err := repo.NewRepository(".")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Create a tree from the current index
	tree := &objects.Tree{Entries: make([]objects.TreeEntry, 0)}
	idx, err := index.ReadIndex(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading index: %v\n", err)
		os.Exit(1)
	}

	for path, entry := range idx.Entries {
		tree.Entries = append(tree.Entries, objects.TreeEntry{
			Mode: entry.Mode,
			Name: path,
			Hash: entry.Hash,
		})
	}

	// Write tree object
	treeData := tree.Serialize()
	treeHash := tree.Hash()
	err = r.WriteObject(treeData, treeHash)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing tree: %v\n", err)
		os.Exit(1)
	}

	// Create commit object
	refs := refs.NewRefs(r)
	parentHash, _ := refs.ReadHEAD()

	commit := &objects.Commit{
		Tree:      treeHash,
		Parents:   []string{},
		Author:    "User <user@example.com>",
		Committer: "User <user@example.com>",
		Message:   message,
		Timestamp: time.Now(),
	}

	if parentHash != "" {
		commit.Parents = append(commit.Parents, parentHash)
	}

	// Write commit object
	commitData := commit.Serialize()
	commitHash := commit.Hash()
	err = r.WriteObject(commitData, commitHash)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing commit: %v\n", err)
		os.Exit(1)
	}

	// Update HEAD
	err = refs.WriteRef("refs/heads/main", commitHash)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error updating HEAD: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created commit %s\n", commitHash)
}
