package main

import(
	"github.com/bytesundso/ScanMC/pkg/scanner"
	"fmt"
)

func main() {
	results, err := scanner.Scan("hosts.txt", 25565, 10000, 1000000000)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(results.Scanned)
}