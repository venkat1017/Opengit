package commands

import (
	"Opengit/repo"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CheckIgnore(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: Opengit check-ignore <path>...")
		os.Exit(1)
	}

	r, err := repo.NewRepository(".")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	ignorePatterns, err := loadGitignore(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading .gitignore: %v\n", err)
		os.Exit(1)
	}

	for _, path := range args {
		relPath, _ := filepath.Rel(r.Worktree, path)
		if isIgnored(relPath, ignorePatterns) {
			fmt.Println(path)
		}
	}
}

func loadGitignore(r *repo.Repository) ([]string, error) {
	file, err := os.Open(r.GitFile("../.gitignore"))
	if os.IsNotExist(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	defer file.Close()

	var patterns []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") {
			patterns = append(patterns, line)
		}
	}
	return patterns, scanner.Err()
}

func isIgnored(path string, patterns []string) bool {
	for _, pattern := range patterns {
		if matched, _ := filepath.Match(pattern, path); matched {
			return true
		}
	}
	return false
}
