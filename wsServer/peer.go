package wsServer

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"vrcdb/httpServer/helpers"
	"vrcdb/wsServer/messages"

	"github.com/gorilla/websocket"
)

type QueuedMessage struct {
	Type int
	Data []byte
}
type Peer struct {
	authValues *helpers.ContextAuthValues
	Connection *websocket.Conn
	WriteChan  chan QueuedMessage
}

func CreatePeer(w http.ResponseWriter, r *http.Request) (*Peer, error) {
	connection, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	peer := &Peer{
		authValues: helpers.GetContextAuthValues(r.Context()),
		Connection: connection,
		WriteChan:  make(chan QueuedMessage, sendChanBufferSize),
	}

	globPeers.Store(peer.authValues.ClientKey, peer)

	return peer, nil
}
func (a *Peer) yeet() error {
	globPeers.Delete(a.authValues.ClientKey)
	return a.Connection.Close()
}
func (a *Peer) writePump() {
	for {
		msg, ok := <-a.WriteChan
		if !ok {
			return
		}

		err := a.Connection.WriteMessage(msg.Type, msg.Data)
		if err != nil {
			log.Print(err)
			return
		}
	}
}
func (a *Peer) readPump() {
	defer close(a.WriteChan) // Will clone writePump when this closes

	for {
		msgType, reader, err := a.Connection.NextReader()
		if err != nil {
			log.Print(err)
			return
		}

		if !a.handleMessage(msgType, reader) {
			return
		}
	}
}
func (a *Peer) handleMessage(msgType int, reader io.Reader) bool {
	switch msgType {
	case websocket.TextMessage:
		return a.handleTextMessage(reader)
	case websocket.BinaryMessage:
		return a.handleBinaryMessage(reader)
	case websocket.PingMessage:
		return a.handlePingMessage(reader)
	case websocket.CloseMessage:
		log.Println("Closing websocket...")
	}

	return false
}
func (a *Peer) handleTextMessage(reader io.Reader) bool {
	var message messages.Message
	err := json.NewDecoder(reader).Decode(&message)

	if err != nil {
		log.Printf("Peer.handleTextMessage() Decode: %v\n", err)
		return false
	}

	switch message.Id {
	case messages.SessionDirectId:
		return a.handleMessageSessionDirect(message.Data)
	}

	return false
}
func (a *Peer) handleBinaryMessage(reader io.Reader) bool {
	// DODO implement me
	return false
}
func (a *Peer) handlePingMessage(reader io.Reader) bool {
	a.WriteChan <- QueuedMessage{
		Type: websocket.PongMessage,
		Data: nil,
	}

	log.Println("Got ping!")

	return true
}

func (a *Peer) handleMessageSessionDirect(data interface{}) bool {
	// TODO check if they are in a session, if they are send it to partner

	// Simulate that:

	bytes, err := json.Marshal(data)
	if err != nil {
		log.Print(err)
		return false
	}

	msg := QueuedMessage{
		Type: websocket.TextMessage,
		Data: bytes,
	}

	globPeers.Range(func(key, value interface{}) bool {
		value.(*Peer).WriteChan <- msg
		return true
	})

	return true
}
