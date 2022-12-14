package mc

import(
	"github.com/xrjr/mcutils/pkg/ping"
	"errors"
	"time"
)

func Ping(hostname string, port int, timeout time.Duration) (ping.JSON, int, error) {
	client := ping.NewClient(hostname, port)
	client.DialTimeout = timeout
	client.ReadTimeout = timeout

	err := client.Connect()
	if err != nil {
		return nil, -1, err
	}

	hsk, err := client.Handshake()
	if err != nil {
		return nil, -1, err
	}

	lat, err := client.Ping()

	if err != nil && !errors.Is(err, ping.ErrInvalidPacketType) {
		return nil, -1, err
	}

	err = client.Disconnect()
	if err != nil {
		return nil, -1, err
	}

	return hsk.Properties, lat, nil
}