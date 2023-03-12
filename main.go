package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	version := getVersion()

	if len(args) == 0 {
		fmt.Println("Please provide an argument")
		fmt.Println("Usage: rmnode [list -ls --list|del -d --del]")
		return
	}

	switch args[0] {
	case "list", "-ls", "--list":
		listAllNodeModules()
	case "del", "-d", "--del":
		deleteAllNodeModules()
	case "help", "-h", "--help":
		fmt.Println("rmnode version", version)
		fmt.Println("Usage: rmnode [list -ls --list|del -d --del]")
	default:
		fmt.Println("Please provide a valid argument")
		fmt.Println("Usage: rmnode [list -ls --list|del -d --del]")
	}
}

func getVersion() string {
	return "1.0.1"
}
