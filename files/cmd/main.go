package main

import (
	"files"
	"log"
	"os"
)

func main() {
	posts, err := files.NewPostsFromFS(os.DirFS("../posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(posts)
}
