package main

import(
	"github.com/bytesundso/ScanMC/pkg/scanner"
	"github.com/bytesundso/ScanMC/internal/db"
	"github.com/bytesundso/ScanMC/internal/file"
	"log"
)

func main() {
	dbcon, err := db.Connect("")
	if err != nil {
		log.Fatal(err)
	}

	file, err := file.LoadFile("hosts.txt")
	if err != nil {
		log.Fatal(err)
	}

	results := make(chan *scanner.ServerInfo)
	go scanner.Scan(file, 25565, results, 10000, 1000000000)
	for true {
		r := <- results
		db.Add[scanner.ServerInfo](dbcon, r)
	}

	db.Close(dbcon)
}