package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
	"time"

	"github.com/saintfish/chardet"

	"golang.org/x/text/transform"

	"golang.org/x/text/encoding/charmap"

	"github.com/asticode/go-astisub"
)

type marker struct {
	StartAt  time.Duration
	Lines    []string
	Duration time.Duration
}

func main() {
	var filePath string
	switch len(os.Args) {
	case 2:
		filePath = os.Args[1]
	case 1:
		filePath = os.Args[0]
	default:
		fmt.Println("Wrong arguments! Example: subkers <PATH TO SUBTITLES>")
		os.Exit(1)
	}
	pathToMarkers, err := ProcessSub(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Success!")
	fmt.Println("Path to markers:", pathToMarkers)
}

// ProcessSub process any subtitles to .csv format (markers) for Adobe Audition and returns path to file
func ProcessSub(filePath string) (string, error) {
	rawMarkers, err := markersFromFile(filePath)
	if err != nil {
		return "", err
	}
	markersFile, err := os.Create(strings.Join(strings.Split(filePath, ".")[:1], "") + ".csv")
	if err != nil {
		return "", errors.New(fmt.Sprint("Can't create file:", err))
	}
	defer markersFile.Close()
	if _, err := markersFile.WriteString("Name\tStart\tDuration\tTime Format\tType\tDescription\n"); err != nil {
		return "", errors.New(fmt.Sprint("Can't write to file:", err))
	}
	for _, val := range rawMarkers {
		if strings.ReplaceAll(strings.Join(val.Lines, " "), "\\N", "") == "" {
			continue
		}
		if _, err := markersFile.WriteString(strings.ReplaceAll(strings.Join(val.Lines, " "), "\\N", "") + "\t" + timeToString(val.StartAt) + "\t" + timeToString(val.Duration) + "\tdecimal\tCue\t\n"); err != nil {
			return "", errors.New(fmt.Sprint("Can't write to file:", err))
		}
	}
	return strings.Join(strings.Split(filePath, ".")[:1], "") + ".csv", nil
}

func markersFromFile(filePath string) ([]marker, error) {
	subs, err := astisub.OpenFile(filePath)
	if err != nil {
		return nil, errors.New(fmt.Sprint("Wrong file format:", err))
	}
	subs.Optimize()
	subs.Unfragment()
	subs.Order()
	return doMarkers(subs)
}

func doMarkers(subs *astisub.Subtitles) ([]marker, error) {
	// Distill subtitles to markers
	var markers []marker
	for _, item := range subs.Items {
		var m marker
		m.StartAt = item.StartAt
		m.Duration = item.EndAt - item.StartAt
		for _, line := range item.Lines {
			for _, lineItem := range line.Items {
				m.Lines = append(m.Lines, string(lineItem.Text))
			}
		}
		markers = append(markers, m)
	}

	result, err := decodeMarkers(markers)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func decodeMarkers(markers []marker) ([]marker, error) {
	var wholeFile []byte
	for _, m := range markers {
		wholeFile = append(wholeFile, []byte(strings.Join(m.Lines, ""))...)
	}

	// Find out the encoding
	detector := chardet.NewTextDetector()
	result, err := detector.DetectBest(wholeFile)
	if err != nil {
		return nil, errors.New(fmt.Sprint("Detector error:", err))
	}
	// Process based on encoding
	switch result.Charset {
	case "UTF-8":
		// just let it go
		return markers, nil
	default:
		// fuck it, let it be so
		for i, m := range markers {
			for j, l := range m.Lines {
				I := bytes.NewReader([]byte(l))
				O := transform.NewReader(I, charmap.Windows1251.NewDecoder())
				line, e := ioutil.ReadAll(O)
				if e != nil {
					return nil, errors.New(fmt.Sprint("Can't read:", err))
				}
				markers[i].Lines[j] = string(line)
			}
		}
		return markers, nil
	}
}

func timeToString(t time.Duration) string {
	ms := fmt.Sprintf("%3.f", math.Floor(math.Mod(t.Seconds(), 1)*1000))
	sec := fmt.Sprintf("%2.f", math.Floor(math.Mod(t.Seconds(), 60)))
	min := fmt.Sprintf("%2.f", math.Floor(t.Seconds()/60))
	return strings.ReplaceAll(min+":"+sec+"."+ms, " ", "")
}
