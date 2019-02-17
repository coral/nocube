package ws

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/coral/nocube/pkg"
	"github.com/gorilla/websocket"
)

type WS struct {
	connections []*websocket.Conn
}

func New() *WS {
	return &WS{}
}

func (ws *WS) Init() {
	http.HandleFunc("/stream", ws.initWS)
	http.ListenAndServe("0.0.0.0:9000", nil)
}

func (ws *WS) ModuleName() string {
	return "WS"
}
func (ws *WS) Write(d []pkg.ColorLookupResult) {

	for _, c := range ws.connections {
		c.WriteJSON(d)
	}
}

func (ws *WS) initWS(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	log.WithFields(log.Fields{
		"IP": r.RemoteAddr,
	}).Info("WS Stream connected")
	ws.connections = append(ws.connections, c)

	go ws.connHandler(c)

}

func (ws *WS) connHandler(conn *websocket.Conn) {

	defer conn.Close()

	for {
		messageType, _, err := conn.ReadMessage()

		//TODO Handle messagetype
		_ = messageType
		if err != nil {
			log.Info("Disconnected WS stream")
			for i, rC := range ws.connections {
				if rC == conn {
					ws.connections = append(ws.connections[:i], ws.connections[i+1:]...)
				}
			}
			break
		}
	}
}
