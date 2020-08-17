package opc

import (
	opc "github.com/coral/go-opc"
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/utils"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

type OPC struct {
	targetFrameRate int
	c               *opc.Client
	connected       bool
}

type writeBuffer struct {
	conn   *websocket.Conn
	buffer chan []pkg.Pixel
}

func New() *OPC {
	return &OPC{
		targetFrameRate: 100,
	}
}

func (ws *OPC) Init() {
	ws.c = opc.NewClient()

	ws.connected = false
	err := ws.c.Connect("tcp", "localhost:7890")
	if err != nil {
		log.WithField("error", err).Error("Could not connect to OPC simulator")
	} else {
		ws.connected = true
	}
}

func (ws *OPC) ModuleName() string {
	return "OPC"
}

func (ws *OPC) SetTargetFrameRate(i int) {
	ws.targetFrameRate = i
}

func (ws *OPC) GetTargetFrameRate() int {
	return ws.targetFrameRate
}

func (ws *OPC) Write(d []pkg.Pixel) {
	if ws.connected {
		m := opc.NewMessage(0)
		m.SetLength(3 * uint16(len(d)))

		for i, p := range d {
			m.SetPixelColor(
				i,
				utils.Clamp255(p.Color[0]*255),
				utils.Clamp255(p.Color[1]*255),
				utils.Clamp255(p.Color[2]*255),
			)
		}

		ws.c.Send(m)
	}

}
