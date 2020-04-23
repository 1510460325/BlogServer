package fileUtil

import (
	"bufio"
	"io"
	"os"
)

func ReadFile(path string) (result string) {
	f, _ := os.Open(path)
	defer func() { _ = f.Close() }()
	bufReader := bufio.NewReader(f)
	already := make([]byte, 0, 1024)
	buf := make([]byte, 1024)
	for {
		readNum, err := bufReader.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == readNum {
			break
		}
		already = append(already, buf[0: readNum]...)
	}
	return string(already)
}
