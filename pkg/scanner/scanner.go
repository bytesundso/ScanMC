package scanner

import(
	"github.com/bytesundso/ScanMC/internal/file"
	"time"
)

type ScanResults struct {
	scanned int64
	found int64
	time time.Time
}

func Scan(filename string, port int, threads int) (ScanResults, error) {
	ips, err := file.LoadFile(filename)
	if err != nil {
		return ScanResults { 0, 0, time.Time { } }, err
	}

	ips.mutex.

	return ScanResults { 0, 0, time.Time { } }, nil
}