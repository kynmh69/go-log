package main

import (
	"io"
	"kynmh69/logging/logging"
	"os"
)

func main() {
	filePath := "./log/"
	fp := os.OpenFile()
	// Logger initilize.
	log := logging.New([]io.Writer{os.Stdout})
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
	log.Critical("eee")
}
