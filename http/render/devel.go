package render

import (
	"context"
	"github.com/fsnotify/fsnotify"
	"io/fs"
	"log"
	"path/filepath"
	"strings"
)

// RebuildOnChange is a render option which will trigger the rebuild of the templates
// when a .html file within baseDir or any subdirectories are changed.  If baseDir is
// the empty string, this will do nothing.
func RebuildOnChange(ctx context.Context, baseDir string) ConfigOption {
	if baseDir == "" {
		return func(config *Config) {}
	}

	return func(cfg *Config) {
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

			if err := filepath.WalkDir(baseDir, func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}
				info, err := d.Info()
				if err != nil {
					return err
				}

				if info.IsDir() {
					if err := watcher.Add(path); err != nil {
						log.Println(err)
					}
				}
				return nil
			}); err != nil {
				log.Println(err)
				return
			}

			<-done
		}()
	}
}
