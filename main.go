package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
)

const version = "0.0.1"

type application struct {
	data any
}

func run() error {
	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "japi is a tool for creating an API from a json file\n\n")
		fmt.Fprintf(os.Stdout, "Usage:\n")
		fmt.Fprintf(os.Stdout, "\tjapi [options]\n\n")

		flag.PrintDefaults()
	}

	file := flag.String("file", "", "json file representing api")
	port := flag.Int("port", 3001, "server port")
	printVersion := flag.Bool("version", false, "print version and exit")
	flag.Parse()

	if *printVersion {
		fmt.Printf("japi version %s\n", version)
		os.Exit(0)
	}

	if *file == "" {
		return errors.New("no file given")
	}

	data, err := jsonFromFile(*file)
	if err != nil {
		return err
	}

	app := &application{
		data: data,
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf("localhost:%d", *port),
		Handler: app.routes(),
	}

	fmt.Fprintf(os.Stdout, "server starting at %s\n", srv.Addr)
	return srv.ListenAndServe()
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		fmt.Fprintf(os.Stderr, "Run 'japi help' for usage.\n")
		os.Exit(1)
	}
}
