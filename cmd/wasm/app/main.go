package main

import (
	"bytes"
	"log"
	"strings"
	"syscall/js"

	"github.com/Terisback/subkers"
	"github.com/maxence-charriere/go-app/v6/pkg/app"
)

type Subtitle struct {
	app.Compo
	filename string
	content  string
}

func (h *Subtitle) Render() app.UI {
	return app.Div().ID("subtitle").Body(
		app.Span().ID("filename").Body(
			app.Text(h.filename),
		),
		app.Div().Class("download").Body(
			app.I().ID("icon").Class("fas fa-download fa-lg"),
		).OnClick(h.OnDownload),
	)
}

func (h *Subtitle) OnDownload(src app.Value, e app.Event) {
	app.Window().Call("onDownload", h.filename, h.content)
	h.Update()
}

var subtits SubtitleList

type SubtitleList struct {
	app.Compo
	subs []app.Node
}

func (h *SubtitleList) Render() app.UI {
	return app.Div().Class("scrollbar style-srl subtitles").Body(
		h.subs...,
	)
}

type Application struct {
	app.Compo
}

func (h *Application) Render() app.UI {
	return app.Div().Body(
		app.Main().Body(
			app.Div().ID("header").Body(
				app.Div().ID("aspect-ratio").Body(
					app.Div().ID("centered").Body(
						app.Img().ID("logo").Src("/web/subkers.png"),
						app.Span().ID("title").Body(app.Text("Subkers")),
					),
				),
			),
			app.Div().ID("application").Body(
				app.Div().ID("header").Body(
					app.Span().Body(
						app.Text("Converted"),
					),
					app.Input().ID("file").Class("inputfile").Name("file").
						OnChange(h.OnInputChange).Type("file").Multiple(true).
						Accept(".srt, .ass, .ssa, .stl, .vtt, .ttml"),
					app.Label().For("file").Body(
						app.I().Class("fas fa-plus fa-lg"),
					),
				),
				&subtits,
			),
		),
	)
}

func (h *Application) OnInputChange(src app.Value, e app.Event) {
	app.Window().Call("fileFromInput")
	h.Update()
}

func Process(this js.Value, inputs []js.Value) interface{} {
	name := inputs[0].String()
	text := inputs[1].String()
	n := strings.Split(name, ".")
	t, err := subkers.SubtitlesType(n[len(n)-1])
	if err != nil {
		log.Println(err)
		return nil
	}
	m, err := subkers.ProcessSpecific(t, strings.NewReader(text))
	if err != nil {
		log.Println(err)
		return nil
	}

	var out bytes.Buffer
	subkers.WriteAll(m, &out)
	what := out.String()
	subtits.subs = append(subtits.subs,
		&Subtitle{filename: strings.Join(n[:len(n)-1], "") + ".csv", content: what})
	subtits.Update()
	return js.ValueOf(what)
}

func main() {
	js.Global().Set("process", js.FuncOf(Process))
	app.Route("/", &Application{})
	app.Run()
}
