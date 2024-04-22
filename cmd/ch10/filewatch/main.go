package main

import (
	"fmt"

	"gopkg.in/fsnotify.v1"
)

func main() {
	counter := 0
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				fmt.Println("event:", event)
				if event.Op&fsnotify.Create == fsnotify.Create {
					fmt.Println("created file:", event.Name)
					counter++
				} else if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("modified file:", event.Name)
					counter++
				} else if event.Op&fsnotify.Remove == fsnotify.Remove {
					fmt.Println("removed file:", event.Name)
					counter++
				} else if event.Op&fsnotify.Rename == fsnotify.Rename {
					fmt.Println("renamed file:", event.Name)
					counter++
				} else if event.Op&fsnotify.Chmod == fsnotify.Chmod {
					fmt.Println("chmod file:", event.Name)
					counter++
				}
			case err := <-watcher.Errors:
				fmt.Println("error:", err)
			}
			if counter > 3 {
				done <- true
			}
		}
	}()

	err = watcher.Add(".")
	if err != nil {
		fmt.Println(err)
	}
	<-done
}
