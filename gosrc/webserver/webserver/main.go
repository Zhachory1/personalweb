package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"

	"webserver/handlers"
)

var (
	port = flag.Int("port", 8080, "port to server this web server on")
	dir  = flag.String("dir", "", "directory that contains all the template files")
	t    *template.Template
	err  error
)

//TODO(zhach): Handle multiple connections
//TODO(zhach): Allow for more expressive templating
//TODO(zhach): Intregrate stubby calls into another server
//TODO(zhach): Make worker/job allocation for connections
//TODO(zhach): 404 page, you idiot
func Init() {
	// Parsing static files into a template.
	handlers.Init(*dir)
	http.HandleFunc("/", handlers.HandleIndex)
}

func main() {
	flag.Parse()
	if *dir == "" {
		panic("--dir flag cannot be empty.")
	}
	Init()
	fmt.Print(
		fmt.Sprintf(
			"\n\nWeb server now serving on http://localhost:%v...\n\n", *port))
	if err := http.ListenAndServe(fmt.Sprintf(":%v", *port), nil); err != nil {
		panic(err)
	}
}
