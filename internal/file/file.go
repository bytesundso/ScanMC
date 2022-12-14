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
	len int
	pos int
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

	return &File { lines, sync.Mutex { }, len(lines), 0 }, nil
}

func ReadNextLine(file *File) (*string, error) {
	file.mutex.Lock()
	defer file.mutex.Unlock()
	
	if(file.pos >= len(file.lines)) {
		return nil, io.EOF
	}

	line := &file.lines[file.pos]
	file.pos++

	percent := float64(file.pos) / float64(file.len)

	fmt.Printf("\033[0;0H")
	fmt.Printf("Progress: %f\n", percent)

	return line, nil
}