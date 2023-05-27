package main

import (
	"io"
	"kynmh69/logging/logging"
	"log"
	"os"
	"time"
)

func main() {
	fp, err := logging.CreateFileStream("")
	if err != nil {
		log.Fatalln("Cannot open file.", err.Error())
	}
	// Logger initilize.
	log := logging.New([]io.Writer{os.Stdout, fp})
	// set log level
	level := logging.DEBUG
	// open env file
	// test log
	log.LogLevel = level
	log.Debug("aaa")
	log.Info("bbb")
	log.Warning("ccc")
	log.Error("ddd")
	time.Sleep(5 * time.Second)
	log.Critical("eee")
}
