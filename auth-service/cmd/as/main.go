package main

import (
	"fmt"

	"github.com/SoulStalker/edreports-auth-servive/internal/config"
)

func main() {
	// init config 
	cfg := config.MustLoad()

	fmt.Println(cfg)
	// init logger

	// init app

	// start grpc server
}