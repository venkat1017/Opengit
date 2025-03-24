package main

import (
	"Opengit/commands"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: Opengit <command> [<args>]")
		fmt.Println("\nAvailable commands:")
		fmt.Println("  init          Create an empty Opengit repository")
		fmt.Println("  add           Add file contents to the index")
		fmt.Println("  rm            Remove files from the working tree and index")
		fmt.Println("  commit        Record changes to the repository")
		fmt.Println("  checkout      Switch branches or restore working tree files")
		fmt.Println("  hash-object   Compute object ID and optionally creates a blob from a file")
		fmt.Println("  cat-file      Provide content of repository objects")
		fmt.Println("  ls-tree       List the contents of a tree object")
		fmt.Println("  ls-files      List files in the index")
		fmt.Println("  status        Show the working tree status")
		fmt.Println("  log           Show commit logs")
		fmt.Println("  tag           Create a tag object")
		fmt.Println("  show-ref      List references")
		fmt.Println("  rev-parse     Parse revisions")
		fmt.Println("  check-ignore  Check if files are ignored by .gitignore")
		os.Exit(1)
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	switch cmd {
	case "init":
		commands.Init(args)
	case "add":
		commands.Add(args)
	case "rm":
		commands.Rm(args)
	case "commit":
		commands.Commit(args)
	case "checkout":
		commands.Checkout(args)
	case "hash-object":
		commands.HashObject(args)
	case "cat-file":
		commands.CatFile(args)
	case "ls-tree":
		commands.LsTree(args)
	case "ls-files":
		commands.LsFiles(args)
	case "status":
		commands.Status(args)
	case "log":
		commands.Log(args)
	case "tag":
		commands.Tag(args)
	case "show-ref":
		commands.ShowRef(args)
	case "rev-parse":
		commands.RevParse(args)
	case "check-ignore":
		commands.CheckIgnore(args)
	default:
		fmt.Printf("Unknown command: %s\n", cmd)
		os.Exit(1)
	}
}
