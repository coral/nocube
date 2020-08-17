package ws

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/coral/nocube/pkg"
	"github.com/gorilla/websocket"
)

type WS struct {
	writeBuffers    []writeBuffer
	targetFrameRate int
}

type writeBuffer struct {
	conn   *websocket.Conn
	buffer chan []pkg.Pixel
}

func New() *WS {
	return &WS{
		targetFrameRate: 30,
	}
}

func (ws *WS) Init() {
	http.HandleFunc("/stream", ws.initWS)
	http.ListenAndServe("0.0.0.0:9000", nil)
}

func (ws *WS) ModuleName() string {
	return "WS"
}

func (ws *WS) SetTargetFrameRate(i int) {
	ws.targetFrameRate = i
}

func (ws *WS) GetTargetFrameRate() int {
	return ws.targetFrameRate
}

func (ws *WS) Write(d []pkg.Pixel) {

	for _, c := range ws.writeBuffers {

		select {
		case c.buffer <- d:
		default:
		}
	}
}

func (ws *WS) initWS(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{}

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	log.WithFields(log.Fields{
		"IP": r.RemoteAddr,
	}).Info("WS Stream connected")

	newBuffer := writeBuffer{
		conn:   c,
		buffer: make(chan []pkg.Pixel, 1),
	}

	ws.writeBuffers = append(ws.writeBuffers, newBuffer)

	go ws.connHandler(c)
	go ws.writeHandler(&newBuffer)

}

func (ws *WS) connHandler(conn *websocket.Conn) {

	defer conn.Close()

	for {
		messageType, _, err := conn.ReadMessage()

		//TODO Handle messagetype
		_ = messageType
		if err != nil {
			log.Info("Disconnected WS stream")
			for i, rC := range ws.writeBuffers {
				if rC.conn == conn {
					ws.writeBuffers = append(ws.writeBuffers[:i], ws.writeBuffers[i+1:]...)
				}
			}
			break
		}
	}
}

func (ws *WS) writeHandler(wb *writeBuffer) {
	tt := time.Second / time.Duration(ws.targetFrameRate)
	ticker := time.NewTicker(tt)

	go func() {
		for _ = range ticker.C {
			d := <-wb.buffer
			wb.conn.WriteJSON(d)
		}
	}()
}
