package main

import (
	"flag"
	"./main"
)

var (
	root = flag.String("directory", ".", "directory to discover")
)

func main() {
	flag.Parse()
	discover.Discover(*root)
}
