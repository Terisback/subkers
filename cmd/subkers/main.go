package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Terisback/subkers"
)

func main() {

	var filePath string

	switch len(os.Args) {
	case 2:
		filePath = os.Args[1]
	default:
		fmt.Println("[ERROR]", "Wrong arguments! Example: subkers <PATH TO SUBTITLES>")
		os.Exit(1)
	}

	splittedFilename := strings.Split(filePath, ".")

	// Getting subs
	markers, err := subkers.Process(filePath)
	if err != nil {
		fmt.Println("[ERROR]", err)
		os.Exit(1)
	}

	// Creating markers file
	markersFile, err := os.Create(strings.Join(splittedFilename[:len(splittedFilename)-1], ".") + ".csv")
	if err != nil {
		fmt.Println(fmt.Sprint("[ERROR] Can't create file: ", err))
		os.Exit(1)
	}
	defer markersFile.Close()

	// Writing markers to file
	if err := subkers.WriteAll(markers, markersFile); err != nil {
		fmt.Println("[ERROR]", err)
		os.Exit(1)
	}

	fmt.Println(fmt.Sprint("[DONE] Successfully converted file: ", filePath))
}
