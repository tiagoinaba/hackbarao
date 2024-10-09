package server

import (
	"github.com/tiagoinaba/hackbarao/database"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
}

func New() *Server {
	return &Server{
		DB: database.New(),
	}
}
