package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func listAllNodeModules() {
	dir := "."

	var wg sync.WaitGroup

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && info.Name() == "node_modules" {
			wg.Add(1)
			go func() {
				fmt.Println(path)
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
}
