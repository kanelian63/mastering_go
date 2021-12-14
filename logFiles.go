package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
	programName := filepath.Base(os.Args[0])
	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)

	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}
	log.Println("LOG_INFO + LOG_LOCAL7: Loggin in GO")

	sysLog, err = syslog.New(sysLog.LOG_MAIL, "some program")
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}
	
	log.Println("LOG_MAIL: logging in go")
	fmt.Println("will u see this?")
}