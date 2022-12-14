package main

import(
	"github.com/bytesundso/ScanMC/pkg/scanner"
	"fmt"
)

func main() {
	results, err := scanner.Scan("/home/david/Schreibtisch/listtool/masscan_cleaned.txt", 25565, 1000, 1000000000)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(results.Scanned)
}