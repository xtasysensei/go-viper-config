package main

import (
	"fmt"
	"log"

	loadcfg "github.com/xtasysensei/go-viper-config/loadCfg"
)

func main() {
	config, err := loadcfg.LoadTomlCfg()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	fmt.Printf("Postgres Host: %s\n", config.Db.Psql.Host)
	fmt.Printf("Postgres Port: %s\n", config.Db.Psql.Port)
	fmt.Printf("Server Port: %s\n", config.SRV.Port)
}
