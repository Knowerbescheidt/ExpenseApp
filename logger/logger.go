package logger

import (
	"fmt"
	"log"
	"os"
)

var WarningLogger *log.Logger
var InfoLogger *log.Logger
var ErrorLogger *log.Logger

func InitLogger() {

	//https://www.youtube.com/watch?v=p45_9nOpD4k minute 11

	if _, err := os.Stat("logs.txt"); err == nil {
		err := os.Remove("logs.txt")
		if err != nil {
			fmt.Println(err)
		}
	}

	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.LstdFlags|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.LstdFlags|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.LstdFlags|log.Lshortfile)
	InfoLogger.Println("Test InfoLogger")
	WarningLogger.Println("Test WarningLogger")
	ErrorLogger.Println("Test ErrorLogger")
	InfoLogger.Println("Loggers have been initialized")
}
