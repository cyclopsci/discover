package discover

import (
	"os"
	"path/filepath"
)

var tree Tree

type Tree struct {
	Dirs	[]string
	Files	[]string
}

func visit(path string, f os.FileInfo, err error) error {
  if f.IsDir() {
	  tree.Dirs = append(tree.Dirs, path)
  } else {
	  tree.Files = append(tree.Files, path)
  }
  return nil
}

func walk(root string) (Tree, error) {
	err := filepath.Walk(root, visit)
	return tree, err
}
