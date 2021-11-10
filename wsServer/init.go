package wsServer

import (
	"flag"
	"fmt"

	"github.com/gorilla/websocket"
)

var (
	Upgrader websocket.Upgrader
	Address  *string
)

func Init(port uint16) {
	Upgrader = websocket.Upgrader{}
	Address = flag.String("addr", fmt.Sprintf("localhost:%v", port), "http service address")
}
