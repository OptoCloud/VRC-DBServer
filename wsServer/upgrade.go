package wsServer

import (
	"log"
	"net/http"
	"sync"
)

const (
	sendChanBufferSize = 64
)

var (
	globPeers sync.Map = sync.Map{}
)

func Upgrade(w http.ResponseWriter, r *http.Request) {
	peer, err := CreatePeer(w, r)
	if err != nil {
		log.Printf("[ERROR] upgrade: %v\n", err)
		return
	}

	go peer.readPump()
	peer.writePump()

	peer.yeet()
}
