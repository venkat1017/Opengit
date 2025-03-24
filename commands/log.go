package commands

import (
	"Opengit/objects"
	"Opengit/refs"
	"Opengit/repo"
	"bytes"
	"fmt"
	"os"
)

func Log(args []string) {
	r, err := repo.NewRepository(".")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	refs := refs.NewRefs(r)
	headHash, err := refs.ReadHEAD()
	if err != nil || headHash == "" {
		fmt.Fprintf(os.Stderr, "No commits yet\n")
		os.Exit(1)
	}

	for hash := headHash; hash != ""; {
		objPath := r.GitFile("objects", hash[:2], hash[2:])
		data, err := os.ReadFile(objPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading commit %s: %v\n", hash, err)
			os.Exit(1)
		}

		content := data[bytes.IndexByte(data, '\x00')+1:]
		commit := objects.ParseCommit(content)
		fmt.Printf("commit %s\n", hash)
		fmt.Printf("Author: %s\n", commit.Author)
		fmt.Printf("Date: %s\n", commit.Timestamp)
		fmt.Printf("\n    %s\n\n", commit.Message)

		if len(commit.Parents) > 0 {
			hash = commit.Parents[0] // Follow first parent
		} else {
			break
		}
	}
}
