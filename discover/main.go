package main

import (
	"fmt"
	"flag"
	"encoding/json"
	"github.com/cyclopsci/discover"
)

var (
	root = flag.String("directory", ".", "directory to discover")
)

func main() {
	flag.Parse()
	results := discover.Run(*root)
	jsonResults, _ := json.MarshalIndent(results, "", "\t")
	fmt.Println(string(jsonResults))
}
