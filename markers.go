package subkers

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/asticode/go-astisub"
)

// Marker structure for Audition
type Marker struct {
	StartAt  time.Duration
	Duration time.Duration
	Lines    []string
}

// markers converts subtitle struct from astisub to []Marker struct
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
	return markers, nil
}

// timeToString converts Marker time variables to text
func timeToString(t time.Duration) string {
	ms := fmt.Sprintf("%3.f", math.Mod(t.Seconds(), 1)*1000)
	sec := fmt.Sprintf("%2.f", math.Floor(math.Mod(t.Seconds(), 60)))
	min := fmt.Sprintf("%.f", math.Floor(t.Minutes()))
	result := min + ":" + sec + "." + ms
	return strings.ReplaceAll(result, " ", "0")
}
