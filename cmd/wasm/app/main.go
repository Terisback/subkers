package main

import (
	"log"
	"strings"
	"syscall/js"

	"github.com/Terisback/subkers"
	"github.com/maxence-charriere/go-app/v6/pkg/app"
)

type hello struct {
	app.Compo
	file string
}

func (h *hello) Render() app.UI {
	return app.Div().Body(
		app.Main().Body(
			app.Input().
				ID("files").
				AutoFocus(true).
				OnChange(h.OnInputChange).
				Type("file").
				Accept(".srt, .ass, .ssa, .stl, .vtt, .ttml"),
			app.Div().ID("list"),
		).ID("application"),
	)
}

func (h *hello) OnInputChange(src app.Value, e app.Event) {
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
	subkers.WriteAll(m, log.Writer())
	return nil
}

func main() {
	js.Global().Set("process", js.FuncOf(Process))
	app.Route("/", &hello{})
	app.Run()
}
