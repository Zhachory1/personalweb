package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/zhach/personalweb/datafiles"
	"github.com/zhach/personalweb/handlers"
)

var (
	port = flag.Int("port", 8080, "port to server this web server on")
)

//TODO(zhach): Handle multiple connections
//TODO(zhach): Allow for more expressive templating
//TODO(zhach): Intregrate stubby calls into another server
//TODO(zhach): Make worker/job allocation for connections
//TODO(zhach): 404 page, you idiot
func Init() {
	// Parsing static files into a template.
	// The datafiles function will return the entire zip file path for
	// the compiled webfiles.
	h := &handlers.Handlers{}
	h.Init(datafiles.GetDataFilesPath("webfiles.zip"))
	http.HandleFunc("/", h.HandleIndex)
}

func main() {
	flag.Parse()
	Init()
	log.Print(
		fmt.Sprintf(
			"\n\nWeb server now serving on http://localhost:%v...\n\n", *port))
	if err := http.ListenAndServe(fmt.Sprintf(":%v", *port), nil); err != nil {
		log.Fatal(err)
	}
}
