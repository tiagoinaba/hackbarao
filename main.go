package main

import "github.com/tiagoinaba/hackbarao/server"

func main() {
	s := server.New()
	s.RegisterRoutes()
	s.Run()
}
