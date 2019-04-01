package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/olahol/melody.v1"
)

type GopherInfo struct {
	ID, X, Y string
}

//web socket
func InitWebsocket(router *gin.Engine) {
	m := melody.New()

	router.GET("/ws/:name", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleConnect(func(session *melody.Session) {
		logrus.Debug("new connetion")
		logrus.Debug(session.IsClosed())
	})

	m.HandleDisconnect(func(session *melody.Session) {
		logrus.Debug("disconnect connection")
		logrus.Debug(session)
		logrus.Debug(session.IsClosed())
	})

	//m.HandleMessage(func(session *melody.Session, bytes []byte) {
	//	logrus.Info("new message")
	//	logrus.Info(string(bytes))
	//	session.Write([]byte("{\"message\":\"ok\"} "))
	//})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.BroadcastFilter(msg, func(q *melody.Session) bool {
			return q.Request.URL.Path == s.Request.URL.Path
		})
	})
}
