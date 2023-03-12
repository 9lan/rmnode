package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args[1:]

	version := getVersion()

	if len(args) == 0 {
		fmt.Println("Please provide an argument")
		fmt.Println("Usage: rmnode [list -ls --list|delete -d --delete]")
		return
	}

	switch args[0] {
	case "list", "-ls", "--list":
		listAllNodeModules()
	case "delete", "-d", "--delete":
		deleteAllNodeModules()
	case "help", "-h", "--help":
		fmt.Println("rmnode version", version)
		fmt.Println("Usage: rmnode [list -ls --list|delete -d --delete]")
	default:
		fmt.Println("Please provide a valid argument")
		fmt.Println("Usage: rmnode [list -ls --list|delete -d --delete]")
	}
}

func listAllNodeModules() {
	root := "."

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && info.Name() == "node_modules" {
			fmt.Println(path)
			return filepath.SkipDir
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
		return
	}
}

func deleteAllNodeModules() {
	root := "."

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && info.Name() == "node_modules" {
			err := os.RemoveAll(path)
			if err != nil {
				return err
			}
			fmt.Printf("Deleted %s\n", path)
			return filepath.SkipDir
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
		return
	}
}

func getVersion() string {
	return "1.0.0"
}
