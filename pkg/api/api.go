package api

import (
	"time"

	"github.com/coral/nocube/pkg/data"

	"github.com/coral/nocube/pkg/api/osc"
	"github.com/coral/nocube/pkg/pipelines"

	"github.com/coral/nocube/pkg/mapping"

	"github.com/coral/nocube/pkg/settings"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	"github.com/sirupsen/logrus"
)

type API struct {
	r *gin.Engine
	m *melody.Melody

	//internal

	mapping   *mapping.Mapping
	pipelines *pipelines.Pipelines
	data      *data.Data
}

func New(m *mapping.Mapping, p *pipelines.Pipelines, d *data.Data) API {
	return API{
		mapping:   m,
		pipelines: p,
		data:      d,
	}
}

func (a *API) Init(s *settings.Settings) {

	o := osc.New(a.pipelines, a.data)
	go o.Init(s.Global.Control.OSC.Listen)
	//gin.SetMode(gin.ReleaseMode)
	a.r = gin.New()
	a.r.Use(cors.Default())
	a.r.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true))
	a.m = melody.New()

	api := a.r.Group("/api")

	////////
	// General & Websocket
	////////

	{
		a.r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong */",
			})
		})

		a.r.GET("/ws", func(c *gin.Context) {
			a.m.HandleRequest(c.Writer, c.Request)
		})

		a.m.HandleMessage(func(s *melody.Session, msg []byte) {
			a.m.Broadcast(msg)
		})

		a.m.HandleConnect(func(s *melody.Session) {
			s.Write([]byte("HELLO HELLO"))
		})
	}

	////////
	// Pipelines
	////////
	{
		api.GET("/pipelines", a.GetPipelines)
	}
	////////
	// Mapping
	////////

	a.r.GET("/mapping", func(c *gin.Context) {
		c.JSON(200, a.mapping.Coordinates)
	})

	type OPCMAP struct {
		Point []float32 `json:"point"`
	}

	a.r.GET("/opcmapping", func(c *gin.Context) {
		ratio := 2.5
		var m []OPCMAP
		for _, e := range a.mapping.Coordinates {
			d := OPCMAP{
				Point: []float32{
					float32(e.Coordinate[0] * ratio),
					float32(e.Coordinate[1] * ratio),
					float32(e.Coordinate[2] * ratio),
				},
			}
			m = append(m, d)
		}
		c.JSON(200, m)
	})

	a.r.Run("0.0.0.0:8000")
}
