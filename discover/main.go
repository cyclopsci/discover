package main

import (
	"fmt"
	"flag"
	"encoding/json"
	"github.com/cyclopsci/discover"
)

var (
	root = flag.String("directory", ".", "directory to discover")
	displayRoot = flag.String("display", ".", "rewrite root directory output to this directory")
)

func main() {
	flag.Parse()
	if *displayRoot == "." {
		*displayRoot = *root
	}
	results := discover.Run(*root, *displayRoot)
	jsonResults, _ := json.MarshalIndent(results, "", "\t")
	fmt.Println(string(jsonResults))
}
