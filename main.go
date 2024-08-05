package main

import (
	"log"
	"net/http"
	"fmt"
	"os"
	"github.com/go-echarts/examples/examples"
)

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {

	examples.LineExamples{}.Examples()
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

// generates only one type of graph - ex kline graph
// func main() {
//     log.Println("Starting server at http://localhost:8090")
//     fmt.Println("Starting server at http://localhost:8090")
//     http.HandleFunc("/", httpServer)
//     err := http.ListenAndServe(":8090", nil)
//     if err != nil {
//         log.Fatalf("Server failed to start: %v", err)
//     }
// }

// func httpServer(w http.ResponseWriter, r *http.Request) {
//     log.Println("Received request")
//     examples.KlineExamples{}.Examples(w)
// }