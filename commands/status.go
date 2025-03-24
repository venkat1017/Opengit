package commands

import (
	"Opengit/index"
	"Opengit/objects"
	"Opengit/refs"
	"Opengit/repo"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// Helper function to check if a tree contains a file with given path and hash
func treeContains(r *repo.Repository, tree *objects.Tree, path string, hash string) bool {
	parts := filepath.SplitList(path)
	current := tree

	// Navigate through directories in the path
	for i := 0; i < len(parts)-1; i++ {
		found := false
		for _, entry := range current.Entries {
			if entry.Name == parts[i] {
				// Read the tree object and continue searching
				treeData, err := r.ReadObject(entry.Hash)
				if err != nil {
					return false
				}
				current = objects.ParseTree(treeData[bytes.IndexByte(treeData, '\x00')+1:])
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	// Check the file itself
	fileName := parts[len(parts)-1]
	for _, entry := range current.Entries {
		if entry.Name == fileName && entry.Hash == hash {
			return true
		}
	}
	return false
}

func Status(args []string) {
	r, err := repo.NewRepository(".")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	refs := refs.NewRefs(r)
	headHash, err := refs.ReadHEAD()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading HEAD: %v\n", err)
		os.Exit(1)
	}

	idx, err := index.ReadIndex(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading index: %v\n", err)
		os.Exit(1)
	}

	// Load HEAD commit’s tree
	var headTree *objects.Tree
	if headHash != "" {
		commitData, _ := r.ReadObject(headHash)
		commit := objects.ParseCommit(commitData[bytes.IndexByte(commitData, '\x00')+1:])
		treeData, _ := r.ReadObject(commit.Tree)
		headTree = objects.ParseTree(treeData[bytes.IndexByte(treeData, '\x00')+1:])
	}

	// Compare HEAD tree and index
	fmt.Println("Changes to be committed:")
	for path, entry := range idx.Entries {
		if headTree == nil || !treeContains(r, headTree, path, entry.Hash) {
			fmt.Printf("  modified:   %s\n", path)
		}
	}

	// Compare index and working tree with concurrency
	fmt.Println("\nChanges not staged for commit:")
	var wg sync.WaitGroup
	changes := make(chan string, 10)
	errChan := make(chan error, 1)

	filepath.Walk(r.Worktree, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || path == r.Gitdir {
			return nil
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			relPath, _ := filepath.Rel(r.Worktree, path)
			data, err := os.ReadFile(path)
			if err != nil {
				errChan <- err
				return
			}
			blob := objects.NewBlob(data)
			hash := blob.Hash()
			if idxEntry, exists := idx.Entries[relPath]; !exists || idxEntry.Hash != hash {
				changes <- fmt.Sprintf("  modified:   %s", relPath)
			}
		}()
		return nil
	})

	go func() {
		wg.Wait()
		close(changes)
		close(errChan)
	}()

	for change := range changes {
		fmt.Println(change)
	}
	if err := <-errChan; err != nil {
		fmt.Fprintf(os.Stderr, "Error scanning working tree: %v\n", err)
		os.Exit(1)
	}
}
