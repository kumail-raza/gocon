package main

import(
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", helloHandler)
	fmt.Println("serving http server on localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request)  {

	log.Println("serving", r.URL)
	fmt.Fprintf(w, "Hello, world!")

}