package main

import (
	"context"
	"fmt"
	"gophercon/bar"
	"gophercon/beach"
	"gophercon/capitol"
	"gophercon/cliffs"
	"gophercon/garden"
	"gophercon/gym"
	"gophercon/jungle"
	"gophercon/port"
	"gophercon/volcano"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	addr := "1234"
	fmt.Printf("Listening on port %s\n", addr)
	log.Fatal(http.ListenAndServe(":"+addr, nil))
}

// Page represents an HTML page.
type Page struct {
	Body template.HTML
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		r.URL.Path = "/home"
	}
	title := strings.TrimPrefix(r.URL.Path, "/")
	if strings.HasPrefix(title, "run") {
		go run(strings.TrimPrefix(title, "run/"))
		return
	}

	p, err := loadPage(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.New("view").Parse(`<div>{{.Body}}</div>`))
	if err := tmpl.Execute(w, p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var cancel context.CancelFunc

func run(title string) {
	if cancel != nil {
		cancel()
	}

	var ctx context.Context
	ctx, cancel = context.WithCancel(context.Background())
	switch title {
	case "bar":
		bar.Bar()
	case "beach":
		beach.Beach()
	case "capitol":
		capitol.Capitol()
	case "cliffs":
		cliffs.Cliffs()
	case "gym":
		gym.Gym()
	case "garden":
		garden.Garden(ctx)
	case "jungle":
		jungle.Jungle()
	case "port":
		port.Port()
	case "volcano":
		volcano.Volcano()
	}
}

func loadPage(title string) (*Page, error) {
	file := title + ".html"
	body, err := ioutil.ReadFile(filepath.Join("static", file))
	if err != nil {
		return nil, err
	}

	body = append(body, fmt.Sprintf(`<script>
	var anHttpRequest = new XMLHttpRequest();
	anHttpRequest.onreadystatechange = function() { 
		if (anHttpRequest.readyState == 4 && anHttpRequest.status == 200)
			aCallback(anHttpRequest.responseText);
	}
	
	anHttpRequest.open( "GET", "http://localhost:1234/run/%s", true );            
	anHttpRequest.send( null );
	</script>
	`, title)...)
	tmp := template.HTML(body)
	return &Page{Body: tmp}, nil
}
