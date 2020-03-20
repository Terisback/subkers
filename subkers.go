package subkers

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/asticode/go-astisub"
)

// SubtitleType needed for ProcessSpecific
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

// SubtitlesType returns SubtitleType from string
func SubtitlesType(ext string) (SubtitleType, error) {

	for i := range subTypes {
		if subTypes[i] == ext {
			return SubtitleType(i), nil
		}
	}

	return SubtitleType(0), errors.New(fmt.Sprint("Can't recognize extension, got", ext, "\nneed", subTypes))
}

// Process subtitles file to .csv format (markers) for Adobe Audition
func Process(filename string) ([]Marker, error) {

	subtitles, err := astisub.OpenFile(filename)
	if err != nil {
		return nil, errors.New(fmt.Sprint("Can't read:", err))
	}

	markers, err := markers(subtitles)
	if err != nil {
		return nil, errors.New(fmt.Sprint("Can't convert to markers:", err))
	}

	decodedMarkers, err := decode(markers)
	if err != nil {
		return nil, errors.New(fmt.Sprint("Can't translate to UTF-8:", err))
	}

	return decodedMarkers, nil
}

// ProcessSpecific process specific extension subtitles to .csv format (markers) for Adobe Audition
func ProcessSpecific(subtitlesType SubtitleType, reader io.Reader) ([]Marker, error) {

	var subtitles *astisub.Subtitles

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
		return nil, errors.New(fmt.Sprint("Wrong file extension, types that can be processed:\n", subTypes))
	}
	if err != nil {
		return nil, errors.New(fmt.Sprint("Can't read:", err))
	}

	markers, err := markers(subtitles)
	if err != nil {
		return nil, errors.New(fmt.Sprint("Can't convert to markers:", err))
	}

	decodedMarkers, err := decode(markers)
	if err != nil {
		return nil, errors.New(fmt.Sprint("Can't translate to UTF-8:", err))
	}

	return decodedMarkers, nil
}

// WriteAll markers to writer
func WriteAll(markers []Marker, writer io.Writer) error {

	// Required line at the beginning of the file
	if _, err := writer.Write([]byte("Name\tStart\tDuration\tTime Format\tType\tDescription\n")); err != nil {
		return errors.New(fmt.Sprint("Can't write:", err))
	}

	// Writing markers
	for _, val := range markers {

		if strings.ReplaceAll(strings.Join(val.Lines, " "), "\\N", "") == "" {
			continue
		}

		// Marker line .. marker line is a cue in the Adobe Audition
		line := strings.ReplaceAll(strings.Join(val.Lines, " "), "\\N", "") + "\t" +
			timeToString(val.StartAt) + "\t" +
			timeToString(val.Duration) + "\tdecimal\tCue\t\n"

		if _, err := writer.Write([]byte(line)); err != nil {
			return errors.New(fmt.Sprint("Can't write:", err))
		}
	}

	return nil
}
