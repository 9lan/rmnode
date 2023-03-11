package main

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestDeleteAllNodeModules(t *testing.T) {
	root := "test"
	err := os.Mkdir(root, 0755)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		err := os.RemoveAll(root)
		if err != nil {
			t.Fatal(err)
		}
	}()

	nodeModules := filepath.Join(root, "node_modules")
	err = os.Mkdir(nodeModules, 0755)
	if err != nil {
		t.Fatal(err)
	}

	_, err = os.Create(filepath.Join(nodeModules, "file.txt"))
	if err != nil {
		t.Fatal(err)
	}

	buf := bytes.NewBufferString("")
	writer := bufio.NewWriter(buf)
	defer writer.Flush()

	os.Args = []string{"rmnode", "-d"}
	deleteAllNodeModules()

	if _, err := os.Stat(nodeModules); !os.IsNotExist(err) {
		t.Errorf("Directory %s was not deleted", nodeModules)
	}
}
