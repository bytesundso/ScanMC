package scanner

import(
	"github.com/bytesundso/ScanMC/pkg/file"
	"github.com/alteamc/minequery/ping"
	//"github.com/xrjr/mcutils/pkg/ping"
	"sync"
	"time"
	"io"
)

type ServerInfo struct {
	Host *string 
	Port int
	Resp *ping.Response
}

func Scan(hosts *file.File, port int, results chan<- *ServerInfo, threads int, timeout time.Duration) {
	wg := sync.WaitGroup { }
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for host, err := file.ReadNextLine(hosts); err != io.EOF; host, err = file.ReadNextLine(hosts) {
				serverInfo, err := pingAddress(host, port, timeout)
				if err == nil {
					results <- serverInfo
				}
			}
		}()
	}

	wg.Wait()
	close(results)
}

func pingAddress(hostname *string, port int, timeout time.Duration) (*ServerInfo, error) {
	res, err := ping.PingWithTimeout(*hostname, uint16(port), timeout)
	if err != nil {
		return nil, err
	}

	return &ServerInfo { hostname, port, res }, nil
}