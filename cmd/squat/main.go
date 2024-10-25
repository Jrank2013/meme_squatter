package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jrank2013/meme_squatter/pkg/web"
	"gopkg.in/yaml.v3"
)

func main() {
	c := flag.String("c", "config.yaml", "path to config file")
	configBytes, err := os.ReadFile(*c)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	var config web.Config
	fmt.Println(string(configBytes))
	yaml.Unmarshal(configBytes, &config)
	fmt.Println(config)

	s, err := web.NewServer(":8080", &config)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := s.Start(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
