package render

import (
	"context"
	"github.com/fsnotify/fsnotify"
	"log"
	"strings"
)

func RebuildOnChange(ctx context.Context, cfg *Config, baseDir string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer watcher.Close()

		done := make(chan bool)
		go func() {
			for {
				select {
				case event, ok := <-watcher.Events:
					if !ok {
						return
					}
					if event.Op&fsnotify.Write == fsnotify.Write {
						if strings.HasSuffix(event.Name, ".html") {
							log.Printf("modified file: %v, rebuilding templates", event.Name)
							cfg.rebuildTemplates()
						}
					}
				case err, ok := <-watcher.Errors:
					if !ok {
						return
					}
					log.Println("error:", err)
				}
			}
		}()

		err = watcher.Add(baseDir)
		if err != nil {
			log.Fatal(err)
		}
		<-done
	}()
}
