package main

import (
	"go-sample/cmd"
	"go-sample/config"
)

func main() {
	cfg := config.GetConfig()
	command := cmd.Initialize(cfg)
	command.Run()
}
