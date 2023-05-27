package main

import "kynmh69/logging/logging"

func main() {
	// Logger initilize.
	log := logging.Logger{}
	log.InitLogger()
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
