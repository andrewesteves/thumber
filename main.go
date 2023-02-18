package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"runtime"
	"time"
)

func main() {
	dir := "./photos"
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	if os.Getenv("WORKERS") == "true" {
		resizeSim(files, dir)
		fmt.Println(time.Since(start))
		return
	}

	resizeSeq(files, dir)
	fmt.Println(time.Since(start))
}

func resizeSim(files []fs.DirEntry, dir string) {
	workers := runtime.GOMAXPROCS(0)
	paths := make(chan string)
	thumbs := make(chan string)

	for worker := 0; worker < workers; worker++ {
		go createThumbnailWorker(paths, thumbs)
	}

	go func() {
		for thumb := range thumbs {
			fmt.Println(thumb)
		}
	}()

	for _, file := range files {
		paths <- fmt.Sprintf("%s/%s", dir, file.Name())
	}
}

func resizeSeq(files []fs.DirEntry, dir string) {
	for _, file := range files {
		path := fmt.Sprintf("%s/%s", dir, file.Name())
		createThumbnail(path)
		fmt.Println(file.Name())
	}
}
