package main 

import (
	"fmt"
	"log"
	"honeypot/internal/config"
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

}
