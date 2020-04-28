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
			"/web/fontawesome.min.js",
			"/web/solid.min.js",
			"/web/FileSaver.js",
		},
		Styles: []string{
			"/web/main.css",
		},
	}

	g := gziphandler.GzipHandler(h)

	if err := http.ListenAndServe(":7777", g); err != nil {
		panic(err)
	}

	// For HTTPS serve
	// if err := http.Serve(autocert.NewListener(), g); err != nil {
	// 	panic(err)
	// }
}

// To build app you need to cd app, the use command GOARCH=wasm GOOS=js go build -o ../app.wasm .
