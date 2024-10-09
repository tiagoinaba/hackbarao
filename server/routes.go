package server

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() {
	s.mux = gin.Default()
	s.mux.LoadHTMLGlob("views/*")
	s.mux.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})
}
