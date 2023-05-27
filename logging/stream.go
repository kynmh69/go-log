package logging

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Create file stream.
// Output log to file stream.
func CreateFileStream(filePath string) (io.Writer, error) {
	if filePath == "" {
		datetime := time.Now()
		fileName := fmt.Sprintf("log-%s", datetime.Format(time.DateOnly))
		filePath = fmt.Sprintf("./log/%s.log", fileName)
	}

	fp, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	return fp, err
}
