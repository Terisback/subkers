package subkers

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/saintfish/chardet"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

// decode needed for translation any to utf-8 (but it does work only with Windows1251)
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
