package web

import (
	"time"

	"github.com/coral/nocube/pkg/mapping"

	"github.com/coral/nocube/pkg/settings"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	"github.com/sirupsen/logrus"
)

type Server struct {
	r *gin.Engine
	m *melody.Melody

	//internal

	mapping *mapping.Mapping
}

func (w *Server) Init(s *settings.Settings, mapping *mapping.Mapping) {
	//gin.SetMode(gin.ReleaseMode)
	w.r = gin.New()
	w.r.Use(cors.Default())
	w.r.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true))
	w.m = melody.New()
	w.mapping = mapping

	w.r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong */",
		})
	})

	w.r.GET("/ws", func(c *gin.Context) {
		w.m.HandleRequest(c.Writer, c.Request)
	})

	w.m.HandleMessage(func(s *melody.Session, msg []byte) {
		w.m.Broadcast(msg)
	})

	w.m.HandleConnect(func(s *melody.Session) {
		s.Write([]byte("HELLO HELLO"))
	})

	w.r.GET("/mapping", func(c *gin.Context) {
		c.JSON(200, w.mapping.Coordinates)
	})

	w.r.Run(s.Global.Control.Web.Listen)
}
