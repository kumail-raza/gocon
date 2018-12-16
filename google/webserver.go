package google

import (
	"fmt"
	"log"
	"net/http"
)

//ServeHTTP ServeHTTP
func ServeHTTP() error {

	http.HandleFunc("/hello", helloHandler)
	fmt.Println("serving http server on localhost:3000")
	return (http.ListenAndServe(":3000", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("serving", r.URL)
	fmt.Fprintf(w, "Hello, world!")

}
