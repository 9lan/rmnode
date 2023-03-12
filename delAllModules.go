package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func deleteAllNodeModules() {
	dir := "."

	var wg sync.WaitGroup

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && info.Name() == "node_modules" {
			wg.Add(1)
			go func() {
				err := os.RemoveAll(path)
				if err == nil {
					fmt.Printf("Deleted %s\n", path)
				}
				wg.Done()
			}()
			return filepath.SkipDir
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	wg.Wait()
	fmt.Println("Done.")
}
