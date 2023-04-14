package main

import (
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mokan-r/snippet-vault/internal/models"
	"log"
	"net/http"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	snippets *models.SnippetModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:123@/snippetvault?parseTime=true", "MySQL data source name")
	flag.Parse()

	infoLog := getInfoLogger()
	errorLog := getErrorLogger()

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()
	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
		snippets: &models.SnippetModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
