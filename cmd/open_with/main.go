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
		fmt.Println("Wrong arguments! Example: subkers <PATH TO SUBTITLES>")
		fmt.Scanln()
		os.Exit(1)
	}

	splittedFilename := strings.Split(filePath, ".")

	// Reading file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Can't read:", err)
		fmt.Scanln()
		os.Exit(1)
	}
	defer file.Close()

	// Recognize file extension
	ext, err := subkers.ExtType(splittedFilename[len(splittedFilename)-1])
	if err != nil {
		fmt.Println(err)
		fmt.Scanln()
		os.Exit(1)
	}

	// Getting subs
	markers, err := subkers.ProcessSpecific(ext, file)
	if err != nil {
		fmt.Println(err)
		fmt.Scanln()
		os.Exit(1)
	}

	// Creating markers file
	markersFile, err := os.Create(strings.Join(splittedFilename[:len(splittedFilename)-1], ".") + ".csv")
	if err != nil {
		fmt.Println(fmt.Sprint("Can't create file:", err))
		fmt.Scanln()
		os.Exit(1)
	}
	defer markersFile.Close()

	// Writing markers to file
	if err := subkers.WriteAll(markers, markersFile); err != nil {
		fmt.Println(err)
		fmt.Scanln()
		os.Exit(1)
	}
}
