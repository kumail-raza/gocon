package main


import(
	google "github.com/minhajuddinkhan/gocon/google"
	"time"
	"fmt"
)

func main() {

	start := time.Now()
	results, err := google.SearchTimeout("golang", 85 * time.Millisecond)
	elapsed := time.Since(start)

	fmt.Println(results)
	fmt.Println(elapsed, err)
}