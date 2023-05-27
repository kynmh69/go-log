package main

import (
	"fmt"
	"io"
	"kynmh69/logging/logging"
	"log"
	"os"
	"time"
)

func main() {
	datetime := time.Now()
	fileName := fmt.Sprintf("log-%s", datetime.Format(time.DateOnly))
	filePath := fmt.Sprintf("./log/%s.log", fileName)

	fp, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660)
	if err != nil {
		log.Fatalln("Cannot open file.", err.Error())
	}
	// Logger initilize.
	log := logging.New([]io.Writer{os.Stdout, fp})
	// set log level
	level := logging.INFO
	// open env file
	// Logger output file and std stream.
	// test log
	log.LogLevel = level
	log.Debug("aaa")
	log.Info("bbb")
	log.Warning("ccc")
	log.Error("ddd")
	time.Sleep(5 * time.Second)
	log.Critical("eee")
}
