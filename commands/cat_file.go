package commands

import (
	"Opengit/objects"
	"Opengit/refs"
	"Opengit/repo"
	"bytes"
	"fmt"
	"os"
)

func CatFile(args []string) {
	if len(args) < 2 || args[0] != "-p" {
		fmt.Println("Usage: Opengit cat-file -p <object>")
		os.Exit(1)
	}

	r, err := repo.NewRepository(".")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	refs := refs.NewRefs(r)
	name := args[1]
	hash := name

	// Handle HEAD and refs
	if name == "HEAD" {
		var err error
		hash, err = refs.ReadHEAD()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading HEAD: %v\n", err)
			os.Exit(1)
		}
	} else if h, err := refs.ReadRef("refs/heads/" + name); err == nil && h != "" {
		hash = h
	} else if h, err := refs.ReadRef("refs/tags/" + name); err == nil && h != "" {
		hash = h
	}

	data, err := r.ReadObject(hash)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading object %s: %v\n", hash, err)
		os.Exit(1)
	}

	// Parse object type and content
	objType := string(data[:bytes.IndexByte(data, ' ')])
	content := data[bytes.IndexByte(data, '\x00')+1:]

	switch objType {
	case "blob":
		fmt.Print(string(content))
	case "tree":
		tree := objects.ParseTree(content)
		for _, entry := range tree.Entries {
			fmt.Printf("%s %s %s\n", entry.Mode, entry.Hash, entry.Name)
		}
	case "commit":
		commit := objects.ParseCommit(content)
		fmt.Printf("tree %s\n", commit.Tree)
		for _, parent := range commit.Parents {
			fmt.Printf("parent %s\n", parent)
		}
		fmt.Printf("author %s\n", commit.Author)
		fmt.Printf("committer %s\n", commit.Committer)
		fmt.Printf("\n%s\n", commit.Message)
	default:
		fmt.Fprintf(os.Stderr, "Unknown object type: %s\n", objType)
		os.Exit(1)
	}
}
