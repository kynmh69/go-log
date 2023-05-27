package logging

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Create file stream.
// Output log to file stream.
// if you wanna output default filepath, please set "" to argument.
func CreateFileStream(filePath string) (*os.File, error) {
	if filePath == "" {
		datetime := time.Now()
		fileName := fmt.Sprintf("log-%s", datetime.Format(time.DateOnly))
		filePath = fmt.Sprintf("./log/%s.log", fileName)
	}

	fp, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	return fp, err
}

func OpenEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Cannot load .env.", err.Error())
	}
}
