package display

import (
	"fmt"
	"io"
	"log"
	"os"
)

var VERBOSE bool
var ANALYSIS bool
var ascii_path = "display/ascii_art.txt"

func AsciiArt() {
	readFile(ascii_path)
}

func Init() {
	AsciiArt()
	if !VERBOSE {
		log.Println("Verbose mode deactivated")
	}
}

func DisplayVerbose(args ...interface{}) {
	if VERBOSE {
		for _, k := range args {
			fmt.Print(k, " ")
		}
		fmt.Println()
	}
}

func DisplayAnalysis(args ...interface{}) {
	if ANALYSIS {
		for _, k := range args {
			fmt.Print(k, " ")
		}
		fmt.Println()
	}
}

func readFile(path string) {
	// Open file for reading.
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()
	// Read file, line by line
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)

		// Break if finally arrived at end of file
		if err == io.EOF {
			break
		}

		// Break if error occured
		if err != nil && err != io.EOF {
			isError(err)
			break
		}
	}
	fmt.Println(string(text))
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}
