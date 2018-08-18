package handlers

import (
	"archive/zip"
	"bytes"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"

	w2b "github.com/zhach/personalweb/w2buffer"
)

type Handlers struct {
	zipFile string
	tmpls   map[string]*template.Template
}

func (h *Handlers) Init(zipFile string) {
	h.zipFile = zipFile
	h.tmpls = make(map[string]*template.Template)

	// Open a zip archive for reading.
	r, err := zip.OpenReader(zipFile)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	for _, f := range r.File {
		log.Printf("Loading template %s...\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		buf := bytes.NewBuffer(nil)
		_, err = io.Copy(buf, rc)
		if err != nil {
			log.Fatal(err)
		}

		tOk := template.New(f.Name)
		template.Must(tOk.Parse(string(buf.Bytes())))
		h.tmpls[f.Name] = tOk

		rc.Close()
		log.Println()
	}
	log.Println("Finished template loading!")
}

func (h *Handlers) Handle404(w http.ResponseWriter) {
	w.Write([]byte("Server incurred an error."))
}

func (h *Handlers) HandleIndex(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	if message == "" {
		message = "Hello World!"
	}
	var templated w2b.WriterToBuffer
	if err := h.tmpls["index.html"].Execute(&templated, message); err != nil {
		h.Handle404(w)
		log.Println(err)
	} else {
		w.Write(templated.GetBuffer())
	}
}
