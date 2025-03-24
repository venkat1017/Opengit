package commands

import (
	"Opengit/refs"
	"Opengit/repo"
	"fmt"
	"os"
	"path/filepath"
)

func ShowRef(args []string) {
	r, err := repo.NewRepository(".")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	refs := refs.NewRefs(r)
	for _, dir := range []string{"refs/heads", "refs/tags"} {
		filepath.Walk(r.GitFile(dir), func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() {
				return nil
			}
			relPath, _ := filepath.Rel(r.Gitdir, path)
			hash, err := refs.ReadRef(relPath)
			if err == nil && hash != "" {
				fmt.Printf("%s %s\n", hash, relPath)
			}
			return nil
		})
	}
}
