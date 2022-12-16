package file

import (
	"sync"
	"os"
	"bufio"
	"io"
	"fmt"
)

type File struct {
	lines []string  
	mutex sync.Mutex
	pos int
	len int
}

func LoadFile(filename string) (*File, error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
    if err != nil {
        return nil, err
    }
	defer f.Close()

	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanLines)

	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return &File { lines, sync.Mutex { }, 0, len(lines) }, nil
}

func ReadNextLine(file *File) (*string, error) {
	file.mutex.Lock()
	defer file.mutex.Unlock()

	if(file.pos >= file.len) {
		return nil, io.EOF
	}

	if file.pos % (file.len / 1000) == 0 {
		fmt.Println(float32(file.pos) / float32(file.len))
	}

	file.pos++
	return &file.lines[file.pos], nil
}