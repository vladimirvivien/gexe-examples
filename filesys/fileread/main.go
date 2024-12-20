package main

import (
	"fmt"
	"os"

	"github.com/vladimirvivien/gexe"
)

// This example shows how to use the fs package to create/write
// files with simplicity.
func main() {
	fmt.Println("Downloading and saving W. E. Du Bois text to /tmp/dubois-souls.txt")

	// Download W. E. Du Bois text from Gutenberg and save locally
	cmd := `wget -O /tmp/dubois-souls.txt  https://www.gutenberg.org/cache/epub/408/pg408.txt`
	if err := gexe.RunProc(cmd).Err(); err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	// read the text display the text
	if err := gexe.FileRead("/tmp/dubois-souls.txt").Into(os.Stdout).Err(); err != nil {
		fmt.Println(err)
	}
}
