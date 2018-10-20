package main

import (
	"fmt"
	"time"

	google "github.com/minhajuddinkhan/gocon/google"
)

func main() {

	start := time.Now()
	results, err := google.Search("golang")
	elapsed := time.Since(start)

	fmt.Println(results)
	fmt.Println(elapsed, err)
}
