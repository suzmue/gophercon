package main

import (
	"context"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/suzmue/gophercon21/bar"
	"github.com/suzmue/gophercon21/beach"
	"github.com/suzmue/gophercon21/garden"
	"github.com/suzmue/gophercon21/gym"
	"github.com/suzmue/gophercon21/jungle"
	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
)

// Page represents an HTML page.
type Page struct {
	Body template.HTML
}

func webServer() {
	http.HandleFunc("/", handler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	addr := "1234"
	fmt.Printf("listening on port %s\n", addr)
	log.Fatal(http.ListenAndServe(":"+addr, nil))
}

var cancel context.CancelFunc

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		r.URL.Path = "/home"
	}
	title := strings.TrimPrefix(r.URL.Path, "/")
	p, err := loadPage(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if cancel != nil {
		cancel()
		cancel = nil
	}

	tmpl := template.Must(template.New("view").Parse(`<div>{{.Body}}</div>`))

	if err := tmpl.Execute(w, p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var ctx context.Context
	ctx, cancel = context.WithCancel(context.Background())
	go func() {
		switch title {
		case "bar":
			bar.Bar(ctx)
		case "gym":
			gym.Gym(ctx)
		case "jungle":
			jungle.Jungle(ctx)
		case "garden":
			garden.Garden(ctx)
		// case "sevenSeas":
		// 	sevenSeas.Sail()
		case "beach":
			beach.Beach(ctx)
		}
	}()
}

func loadPage(title string) (*Page, error) {
	filename := title + ".md"
	body, err := ioutil.ReadFile(filepath.Join("static", filename))
	if err != nil {
		return nil, err
	}

	b, err := renderMarkdown(body)
	if err != nil {
		return nil, err
	}
	return &Page{Body: b}, nil
}

func renderMarkdown(input []byte) (template.HTML, error) {
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM, extension.Table, emoji.Emoji),
		goldmark.WithRendererOptions(html.WithUnsafe()))
	reader := text.NewReader(input)
	doc := md.Parser().Parse(reader)
	var b strings.Builder
	if err := md.Renderer().Render(&b, input, doc); err != nil {
		return "", err
	}
	return template.HTML(b.String()), nil
}
