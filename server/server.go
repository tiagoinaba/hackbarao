package server

import (
	"github.com/gin-gonic/gin"
	"github.com/tiagoinaba/hackbarao/database"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB

	mux *gin.Engine
}

func (s Server) Run() {
	s.mux.Run()
}

func New() *Server {
	return &Server{
		DB: database.New(),
	}
}
