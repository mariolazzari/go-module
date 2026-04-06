package main

import (
	"github.com/tsawler/toolkit"
)

func main() {
	var tools toolkit.Tools

	tools.CreateDirIfNotExist("./test-dir")
}
