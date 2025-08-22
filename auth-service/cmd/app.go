package main

import (
	"log"

	"github.com/SoulStalker/edreports/auth-service/internal/infra"
)

func main() {
	// init db
	cfg := infra.DBConfig{DSN: "host=127.0.0.1 user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Yekaterinburg"}
	_, err := infra.InitDB(cfg)
	if err != nil {
		log.Panicln(err)
	}

}
