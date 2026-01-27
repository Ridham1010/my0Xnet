package discovery

import (
	"log"
	"time"

	"github.com/grandcat/zeroconf"
)

func Advertise(port int, deviceID string) {
	server, err := zeroconf.Register(
		"0Xnet-"+deviceID,
		"_0xnet._tcp",
		"local.",
		port,
		[]string{"deviceId=" + deviceID},
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer server.Shutdown()

	for {
		time.Sleep(time.Second)
	}
}
