package main

import (
	"goinaction/src/gtihub.com/goinaction/code/chapter2/simple/search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
