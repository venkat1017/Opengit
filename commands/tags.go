package commands

import (
	"Opengit/refs"
	"Opengit/repo"
	"fmt"
	"os"
)

func Tag(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: mygit tag <tagname> [<commit-ish>]")
		os.Exit(1)
	}

	r, err := repo.NewRepository(".")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	refs := refs.NewRefs(r)
	tagName := args[0]
	commitHash := ""
	if len(args) > 1 {
		commitHash = args[1]
	} else {
		commitHash, err = refs.ReadHEAD()
		if err != nil || commitHash == "" {
			fmt.Fprintf(os.Stderr, "Error: no commit to tag\n")
			os.Exit(1)
		}
	}

	if !objectExists(r, commitHash) {
		fmt.Fprintf(os.Stderr, "Error: invalid commit %s\n", commitHash)
		os.Exit(1)
	}

	err = refs.WriteRef("refs/tags/"+tagName, commitHash)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing tag: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Created tag %s pointing to %s\n", tagName, commitHash)
}
