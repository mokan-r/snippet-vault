package main

import (
	"log"
	"os"
)

const (
	blueBackground  = "\033[48;5;26m"
	redBackground   = "\033[48;5;124m"
	clearBackground = "\033[0m"
)

func getInfoLogger() *log.Logger {
	return log.New(
		os.Stdout,
		blueBackground+"  INFO "+clearBackground+"\t",
		log.Ldate|log.Ltime,
	)
}

func getErrorLogger() *log.Logger {
	return log.New(os.Stderr,
		redBackground+" ERROR "+clearBackground+"\t",
		log.Ldate|log.Ltime+log.Lshortfile,
	)
}
