package main

import (
	"fmt"
	"flag"
	"./main"
)

var (
	root = flag.String("directory", ".", "directory to discover")
)

func main() {
	flag.Parse()
	results := discover.Discover(*root)
	fmt.Println(results)
}
