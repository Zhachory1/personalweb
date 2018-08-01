package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strings"

	w2b "webserver/w2buffer"
)

var (
	t   *template.Template
	err error
)

func Init(dir string) {
	if t, err = template.ParseGlob(path.Join(dir, "*")); err != nil {
		panic(err)
	}
}

func Handle404(w http.ResponseWriter) {
	w.Write([]byte("Server incurred an error."))
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	if message == "" {
		message = "Hello World!"
	}
	var templated w2b.WriterToBuffer
	if err := t.ExecuteTemplate(&templated, "index", message); err != nil {
		Handle404(w)
		fmt.Print(err)
	} else {
		w.Write(templated.GetBuffer())
	}
}
