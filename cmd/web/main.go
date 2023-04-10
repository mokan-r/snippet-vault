package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

const (
	blueBackground  = "\033[48;5;26m"
	redBackground   = "\033[48;5;124m"
	clearBackground = "\033[0m"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(
		os.Stdout,
		blueBackground+" INFO "+clearBackground+"\t",
		log.Ldate|log.Ltime,
	)
	errorLog := log.New(os.Stderr,
		redBackground+" ERROR "+clearBackground+"\t",
		log.Ldate|log.Ltime+log.Lshortfile,
	)

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})

	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
