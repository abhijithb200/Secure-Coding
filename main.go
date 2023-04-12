package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/fsnotify/fsnotify"
)

func FindHash() string {
	f, err := os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

var counter int

func main() {

	// <-- Need to run as separate process -->
	// fileName := "./layer1/file_create.go"
	// cmd := exec.Command("go", "run", fileName)
	// cmd.CombinedOutput()

	// fileName = "./layer2/diffmain.go"
	// cmd = exec.Command("go", "run", fileName)
	// output, err := cmd.CombinedOutput()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println(string(output))

	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:

				if !ok {
					return
				}
				if event.Has(fsnotify.Write) && event.Name == "parser\\example.go" {

					// the file is changed - because it call two time - restrict to one

					if counter%2 == 0 {
						fileName := "./layer2/diffmain.go"
						cmd := exec.Command("go", "run", fileName)
						output, err := cmd.CombinedOutput()
						if err != nil {
							fmt.Println("Error:", err)
							return
						}
						fmt.Println(string(output))
					}

				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
			counter++
		}
	}()

	// Add a path.
	err = watcher.Add("./parser/example.go")
	if err != nil {
		log.Fatal(err)
	}

	// Block main goroutine forever.
	<-make(chan struct{})

}
