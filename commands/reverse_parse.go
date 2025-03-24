package commands

import (
	"Opengit/refs"
	"Opengit/repo"
	"fmt"
	"os"
)

func RevParse(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: mygit rev-parse <name>")
		os.Exit(1)
	}

	r, err := repo.NewRepository(".")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	refs := refs.NewRefs(r)
	name := args[0]

	// Try as a ref first
	if hash, err := refs.ReadRef("refs/heads/" + name); err == nil && hash != "" {
		fmt.Println(hash)
		return
	}
	if hash, err := refs.ReadRef("refs/tags/" + name); err == nil && hash != "" {
		fmt.Println(hash)
		return
	}
	if hash, err := refs.ReadHEAD(); err == nil && name == "HEAD" && hash != "" {
		fmt.Println(hash)
		return
	}

	// Assume it’s a hash
	if len(name) == 40 && objectExists(r, name) {
		fmt.Println(name)
		return
	}

	fmt.Fprintf(os.Stderr, "Unknown reference or object: %s\n", name)
	os.Exit(1)
}

func objectExists(r *repo.Repository, hash string) bool {
	_, err := r.ReadObject(hash)
	return err == nil
}
