package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/minhajuddinhkhan/gocon/google"
	"github.com/urfave/cli"
)

func main() {

	timeout := 80 * time.Millisecond
	app := cli.NewApp()
	app.Usage = "Application for GDG"
	app.Commands = []cli.Command{
		cli.Command{
			Name:  "serve",
			Usage: "Starts http server on port 3000",
			Action: func(c *cli.Context) error {
				return (google.ServeHTTP())
			},
		},
		cli.Command{
			Name:  "serial",
			Usage: "Fetches result sequentially",
			Action: func(c *cli.Context) error {
				start := time.Now()
				results, err := google.Search("golang")
				elapsed := time.Since(start)
				fmt.Println(results)
				fmt.Println(elapsed, err)
				return nil
			},
		},
		cli.Command{
			Name:  "parallel",
			Usage: "Fetches result in parallel",
			Action: func(c *cli.Context) error {
				start := time.Now()
				results, err := google.SearchParallel("golang")
				elapsed := time.Since(start)
				fmt.Println(results)
				fmt.Println(elapsed, err)

				return nil
			},
		},
		cli.Command{
			Name:  "timeout",
			Usage: "Fetches result in parallel with a timeout",
			Action: func(c *cli.Context) error {
				start := time.Now()
				results, err := google.SearchTimeout("golang", timeout)
				elapsed := time.Since(start)
				fmt.Println(results)
				fmt.Println(elapsed, err)

				return nil
			},
		},
		cli.Command{
			Name:  "replica",
			Usage: "Fetches result from multiple replicated sources in parallel with a timeout",
			Action: func(c *cli.Context) error {
				start := time.Now()
				results, err := google.SearchReplicated("golang", timeout)
				elapsed := time.Since(start)
				fmt.Println(results)
				fmt.Println(elapsed, err)

				return nil
			},
		},
	}

	log.Fatal(app.Run(os.Args))

}
