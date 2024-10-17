package main

import (
	"fmt"
	"os"

	"github.com/jrank2013/meme_squatter/pkg/web"
)

func main() {
	s, err := web.NewServer(":8080")

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	s.Start()
}
