package main

import (
	"net/http"

	"github.com/NYTimes/gziphandler"
	"github.com/maxence-charriere/go-app/v6/pkg/app"
)

func main() {
	h := &app.Handler{
		Title:           "Subkers",
		Name:            "Subkers",
		ShortName:       "Subkers",
		Description:     "A tool for converting subtitles to Audition Markers",
		Author:          "Terisback",
		Version:         "0.2.0",
		LoadingLabel:    "Loading Subkers",
		BackgroundColor: "#212D42",
		ThemeColor:      "#80AEFF",
		Icon: app.Icon{
			Default: "/web/subkers.png",
		},
		Keywords: []string{
			"subkers",
			"subtitles",
			"audition",
			"markers",
			"convertion",
		},
		Scripts: []string{
			"/web/main.js",
		},
		Styles: []string{
			"/web/main.css",
		},
	}

	g := gziphandler.GzipHandler(h)

	if err := http.ListenAndServe(":7777", g); err != nil {
		panic(err)
	}
}
