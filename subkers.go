package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
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

	splittedFilename := strings.Split(filePath, ".")

	// Reading file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Can't read:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Recognize file extension
	ext, err := ExtType(splittedFilename[len(splittedFilename)-1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Getting subs
	markers, err := ProcessSubs(ext, file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Creating markers file
	markersFile, err := os.Create(strings.Join(strings.Split(filePath, ".")[:1], "") + ".csv")
	if err != nil {
		fmt.Println(fmt.Sprint("Can't create file:", err))
		os.Exit(1)
	}
	defer markersFile.Close()

	// Writing markers to file
	if _, err := markersFile.WriteString("Name\tStart\tEnd\tTime Format\tType\tDescription\n"); err != nil {
		fmt.Println(fmt.Sprint("Can't write to file:", err))
		os.Exit(1)
	}
	for _, val := range markers {
		if strings.ReplaceAll(strings.Join(val.Lines, " "), "\\N", "") == "" {
			continue
		}

		line := strings.ReplaceAll(strings.Join(val.Lines, " "), "\\N", "") + "\t" +
			timeToString(val.StartAt) + "\t" +
			timeToString(val.EndAt) + "\tdecimal\tCue\t\n"

		if _, err := markersFile.WriteString(line); err != nil {
			fmt.Println(fmt.Sprint("Can't write to file:", err))
			os.Exit(1)
		}
	}
}

type Marker struct {
	StartAt time.Duration
	Lines   []string
	EndAt   time.Duration
}

type SubtitleType int

const (
	SRT SubtitleType = iota
	SSA
	ASS
	STL
	TTML
	WebVVT
)

var subTypes = []string{
	"srt",
	"ssa",
	"ass",
	"stl",
	"ttml",
	"vtt",
}

// ExtType returns SubtitleType from string
func ExtType(ext string) (SubtitleType, error) {
	for i := range subTypes {
		if subTypes[i] == ext {
			return SubtitleType(i), nil
		}
	}
	return SubtitleType(0), errors.New(fmt.Sprint("Can't recognize extension, got", ext))
}

// ProcessSub process any subtitles to .csv format (markers) for Adobe Audition and returns path to file
func ProcessSubs(subtitlesType SubtitleType, reader io.Reader) ([]Marker, error) {
	var subtitles *astisub.Subtitles
	{
		var err error
		switch subtitlesType {
		case SRT:
			subtitles, err = astisub.ReadFromSRT(reader)
		case SSA, ASS:
			subtitles, err = astisub.ReadFromSSA(reader)
		case STL:
			subtitles, err = astisub.ReadFromSTL(reader)
		case TTML:
			subtitles, err = astisub.ReadFromTTML(reader)
		case WebVVT:
			subtitles, err = astisub.ReadFromWebVTT(reader)
		default:
			return nil, errors.New(fmt.Sprint("Wrong file extension:", err))
		}
		if err != nil {
			return nil, errors.New(fmt.Sprint("Can't read:", err))
		}
	}
	return markers(subtitles)
}

func markers(subs *astisub.Subtitles) ([]Marker, error) {
	subs.Optimize()
	subs.Order()
	subs.Unfragment()

	// Distill subtitles to markers
	var markers []Marker
	for _, item := range subs.Items {
		var m Marker
		m.StartAt = item.StartAt
		m.EndAt = item.EndAt
		for _, line := range item.Lines {
			for _, lineItem := range line.Items {
				m.Lines = append(m.Lines, string(lineItem.Text))
			}
		}
		markers = append(markers, m)
	}

	return decode(markers)
}

func decode(markers []Marker) ([]Marker, error) {
	var buffer []byte
	for _, m := range markers {
		buffer = append(buffer, []byte(strings.Join(m.Lines, ""))...)
	}

	// Find out the encoding
	detector := chardet.NewTextDetector()
	result, err := detector.DetectBest(buffer)
	if err != nil {
		return nil, errors.New(fmt.Sprint("Detector error:", err))
	}

	// Process based on encoding
	// I don't know how to do it multilang
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
