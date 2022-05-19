package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

type countHandler struct {
	mu sync.Mutex //guards n
	n  int
}

func main() {
	handler() //Calls the handler function

	handlerFunc() //Calls the handlerFunc function

	listenAndServe() //Calls the listenAndServe function

	listenAndServeTLS() //Calls the listenAndServerTLS function

} //end of main function

func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
	fmt.Fprintf(w, "count is %d\n", h.n)
}

func handler() { //Registers the handler for the given pattern in the DefaultServeMux^^
	http.Handle("/count", new(countHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlerFunc() { //registers the handler function for the given pattern in the DefaultServeMux
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #2!\n")
	}
	http.HandleFunc("/", h1)
	http.HandleFunc("/endpoint", h2)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func listenAndServe() { //Listens on the TCP network addresss addr and then calls Serve with handler to handle requests on incoming directions
	helloHandler := func(w http.ResponseWriter, req *http.Request) { //Accepted connects are configured to enable TCP keep-alives
		io.WriteString(w, "Hello, world!\n")
	}
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func listenAndServeTLS() { //Same as listenAndServe, but it expects HTTP connections.
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) { //Also expects files containing certificate and matching private key
		io.WriteString(w, "Hello, TLS!\n")
	})
	log.Printf("About to listen on 8443. Go to https://127.0.0.1:8443/")
	err := http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil)
	log.Fatal(err)
}
