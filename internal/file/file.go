package file

import "sync"

type File struct {
	addrs []string
	length int
	position int   
	mutex sync.Mutex
}

func LoadFile(filename string) (File, error) {
	return File { addrs: []string { "asd", "qwe" }, length: 2, position: 0, mutex: sync.Mutex { } }, nil
}

func ReadNextLine(file *File) (string, error) {
	return "line", nil
}