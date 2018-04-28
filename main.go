package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/xlab/treeprint"
)

func main() {
	paths := os.Args[1:]
	tree := treeprint.New()

	for _, path := range paths {
		parent := tree

		for _, element := range strings.Split(filepath.ToSlash(path), "/") {

			existing := parent.FindByValue(element)
			if existing != nil {
				parent = existing
			} else {
				parent = parent.AddBranch(element)
			}

		}
	}

	fmt.Println(tree.String())
}
