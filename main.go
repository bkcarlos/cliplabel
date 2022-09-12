package main

import (
	"github.com/bkcarlos/cliplabel/initialize"
	"github.com/bkcarlos/cliplabel/logger"
)

func init() {
	initialize.Init()
}

func main() {
	logger.Infos("clip lable server is starting")
}
