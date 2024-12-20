package main

import (
	"fmt"
	"os"

	"github.com/vladimirvivien/gexe"
)

// This example shows how to use file operations with gexe.
func main() {
	gexe.SetVar("file", "/tmp/warofworlds.txt")
	fmt.Println(gexe.Eval("Downloading and saving War of the Worlds text: $file"))

	// Download text and save it locally
	cmd := `wget -O $file  https://www.gutenberg.org/cache/epub/36/pg36.txt`
	if err := gexe.RunProc(cmd).Err(); err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	// read the dowloaded file and stream it to stdout
	if err := gexe.FileRead("$file").Into(os.Stdout).Err(); err != nil {
		fmt.Println(err)
	}
}
