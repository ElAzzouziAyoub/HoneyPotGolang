package main 

import (
	"fmt"
	"log"
	"honeypot/internal/config"
	"honeypot/internal/logger"
  "honeypot/internal/server"
  "honeypot/internal/storage"
)

func main(){
	cfg, err := config.Load("config.json")
	if err != nil {
  	  log.Fatal(err)
	}

	if err := cfg.Validate(); err != nil {
  	  log.Fatal(err)
	}

	fmt.Println("config OK")



	logg := logger.NewStdoutLogger()
  store := storage.NewFileStorage("events.jsonl")

  srv := server.New(":2222", logg, store)
  srv.Start()

}
