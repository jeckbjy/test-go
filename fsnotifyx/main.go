package main

import (
	"io/ioutil"
	"log"

	"github.com/fsnotify/fsnotify"
)

func main() {
	log.Printf("run fsnotify")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalln(err)
	}

	watcher.Add(".services")
	for {
		select {
		case event := <-watcher.Events:
			switch {
			case event.Op&fsnotify.Write == fsnotify.Write:
				data, err := ioutil.ReadFile(event.Name)
				log.Printf("Write: %s:%s,%s, %+v", event.Op, event.Name, data, err)
			case event.Op&fsnotify.Create == fsnotify.Create:
				log.Printf("Create: %s: %s", event.Op, event.Name)
			case event.Op&fsnotify.Remove == fsnotify.Remove:
				log.Printf("Remove: %s: %s", event.Op, event.Name)
			case event.Op&fsnotify.Rename == fsnotify.Rename:
				log.Printf("Rename: %s: %s", event.Op, event.Name)
			case event.Op&fsnotify.Chmod == fsnotify.Chmod:
				log.Printf("Chmod:  %s: %s", event.Op, event.Name)
			}
		case err := <-watcher.Errors:
			log.Printf("%+v", err)
		}
	}
}
