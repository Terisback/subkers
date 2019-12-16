package subkers

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"strings"
	"time"

	"github.com/saintfish/chardet"

	"golang.org/x/text/transform"

	"golang.org/x/text/encoding/charmap"

	"github.com/asticode/go-astisub"
)

type Marker struct {
	StartAt  time.Duration
	Duration time.Duration
	Lines    []string
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

// SubType returns SubtitleType from string
func SubType(ext string) (SubtitleType, error) {
	for i := range subTypes {
		if subTypes[i] == ext {
			return SubtitleType(i), nil
		}
	}
	return SubtitleType(0), errors.New(fmt.Sprint("Can't recognize extension, got", ext))
}

// Process subtitles file to .csv format (markers) for Adobe Audition
func Process(filename string) ([]Marker, error) {
	subtitles, err := astisub.OpenFile(filename)
	if err != nil {
		return nil, errors.New(fmt.Sprint("Can't read:", err))
	}
	return markers(subtitles)
}

// ProcessSpecific process specific extension subtitles to .csv format (markers) for Adobe Audition
func ProcessSpecific(subtitlesType SubtitleType, reader io.Reader) ([]Marker, error) {
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

// WriteAll markers to writer
func WriteAll(markers []Marker, writer io.Writer) error {
	if _, err := writer.Write([]byte("Name\tStart\tDuration\tTime Format\tType\tDescription\n")); err != nil {
		return errors.New(fmt.Sprint("Can't write:", err))
	}
	for _, val := range markers {
		if strings.ReplaceAll(strings.Join(val.Lines, " "), "\\N", "") == "" {
			continue
		}

		line := strings.ReplaceAll(strings.Join(val.Lines, " "), "\\N", "") + "\t" +
			timeToString(val.StartAt) + "\t" +
			timeToString(val.Duration) + "\tdecimal\tCue\t\n"

		if _, err := writer.Write([]byte(line)); err != nil {
			return errors.New(fmt.Sprint("Can't write:", err))
		}
	}
	return nil
}

func markers(subs *astisub.Subtitles) ([]Marker, error) {
	// Distill subtitles lines to markers lines
	var markers []Marker
	for _, item := range subs.Items {
		var m Marker
		m.StartAt = item.StartAt
		m.Duration = item.EndAt - item.StartAt
		for _, line := range item.Lines {
			for _, lineItem := range line.Items {
				m.Lines = append(m.Lines, lineItem.Text)
			}
		}
		markers = append(markers, m)
	}

	// return decoded markers in the right encoding
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
	// it will be okay
	switch result.Charset {
	case "UTF-8":
		// its okay
		// just let it go
		return markers, nil
	default:
		// should work almost anyway
		// let it be so
		for i, m := range markers {
			for j, l := range m.Lines {
				O := transform.NewReader(bytes.NewReader([]byte(l)), charmap.Windows1251.NewDecoder())
				line, e := ioutil.ReadAll(O)
				if e != nil {
					return nil, errors.New(fmt.Sprint("Can't read:", err))
				}
				markers[i].Lines[j] = string(line)
			}
		}
		// its okay
		return markers, nil
	}
}

// timeToString converts Marker time variables to text
func timeToString(t time.Duration) string {
	ms := fmt.Sprintf("%3.f", math.Mod(t.Seconds(), 1)*1000)
	sec := fmt.Sprintf("%2.f", math.Floor(math.Mod(t.Seconds(), 60)))
	min := fmt.Sprintf("%.f", math.Floor(t.Minutes()))
	result := min + ":" + sec + "." + ms
	return strings.ReplaceAll(result, " ", "0")
}
