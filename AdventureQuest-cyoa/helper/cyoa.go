package helper

import (
	"encoding/json"
	_ "fmt"
	"html/template"
	"io"
	"net/http"
	"strings"
)

type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type StoryArc struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}

type Story map[string]StoryArc

func ParseJson(f io.Reader) (Story, error) {
	// ``` Function for parsing a json data into struct var ```
	d := json.NewDecoder(f)
	// var story Story
	story := make(Story)
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

var templ = `<!DOCTYPE html>
<html>
  <head>
    <title>{{.Title}}</title>
    <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0-rc.2/css/materialize.min.css">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
  </head>
  <body class="grey lighten-4">
    <div class="row">
        <div class="col m12">
          <div class="card large horizontal z-depth-4 brown white-text">
            <div class="card-stacked">
              <h2 class="card-title">{{.Title}}</h2>
              <div class="card-content">
                {{range .Story}}
                  <p>{{.}}</p>
                {{end}}
              </div>
              <div class="card-action">
                {{range .Options}}
                  <a href="{{.Arc}}">{{.Text}}</p>
                {{end}}
              </div>
            </div>
          </div>
        </div>
      </div>
  </body>
</html>`

func init() {
	tpl = template.Must(template.New("").Parse(templ))
}

var tpl *template.Template

func NewHandler(s Story, t *template.Template) http.Handler {
	if t == nil {
		t = tpl
	}
	return handler{s, t}
}

type handler struct {
	s Story
	t *template.Template
}

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var path string

	if req.URL.Path == "/" {
		path = "intro"
	} else {
		path = strings.TrimLeft(req.URL.Path, "/")
	}
	arc := h.s[path]

	if err := tpl.Execute(w, arc); err != nil {
		panic(err)
	}

}
