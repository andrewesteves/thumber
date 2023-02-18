package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func createThumbnailWorker(paths <-chan string, thumbs chan<- string) {
	for path := range paths {
		createThumbnail(path)
		thumbs <- path
	}
}

func createThumbnail(path string) {
	var args []string
	args = append(args, path)
	args = append(args, "-resize")
	args = append(args, "200x200^")
	args = append(args, "-gravity")
	args = append(args, "center")
	args = append(args, "-extent")
	args = append(args, "200x200")

	paths := strings.Split(path, "/")
	fileName := paths[len(paths)-1]
	paths[len(paths)-1] = fmt.Sprintf("thumb_%s", fileName)
	args = append(args, strings.Join(paths, "/"))

	cmd := exec.Command("convert", args...)
	_, err := cmd.Output()
	if err != nil {
		log.Printf("create thumbnail: %v", err)
	}
}
