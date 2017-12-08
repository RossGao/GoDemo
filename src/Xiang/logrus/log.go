package main

import (
	log "github.com/Sirupsen/logrus"
)

func main() {
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	testValue := "Hello"
	log.SetFormatter(Formatter)
	log.Infof("Some info. Earth is not flat.%v", testValue)
	log.Warning("This is a warning")
	log.Error("Not fatal. An error. Won't stop execution")
	log.Fatal("MAYDAY MAYDAY MAYDAY. Execution will be stopped here")
	log.Panic("Do not panic")
}
