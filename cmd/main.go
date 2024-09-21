package main

import (
	"Odnostranishka/config"
	"Odnostranishka/internal/server"
	lg "gitlab.com/nevasik7/logger"
)

func main() {
	lg.Init()
	server.New(config.MustLoad().Init()).Run()
}
