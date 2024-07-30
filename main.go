package main

import (
	"log"
	"net/http"
	"fmt"
	"os"
	"github.com/go-echarts/examples/examples"
)

// export logRequest
func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
// var data[];
// data[]=getData()


func main() {

	examples.KlineExamples{}.Examples()
	fmt.Println("In the Main Function")

	serverPages := "true"
	if len(os.Args) > 1 {
		serverPages = os.Args[1]
	}

	if serverPages == "false" {
		log.Println("Generated pages only, not server")
		return
	}
	fs := http.FileServer(http.Dir("examples/html"))
	log.Println("running server at http://localhost:8090")
	log.Fatal(http.ListenAndServe("localhost:8090", logRequest(fs)))
}
