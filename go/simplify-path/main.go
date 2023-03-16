package main

import (
	"fmt"
	"path"
)

func main() {
	path := "/home/"
	fmt.Println(simplifyPath(path))
}

func simplifyPath(path1 string) string {
	return path.Clean(path1)
}
