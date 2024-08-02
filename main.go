package main

import (
	"fmt"
	"log"
	"os"
	"time"
	
	"github.com/balagrivine/go_auth/config"
	"github.com/balagrivine/go_auth/handler"
	"github.com/joho/dotenv"
)

func main() {

	// Load environmental variables
	dotenv.Load(".env")

	// Initialize the ocnfiguration
	if apiCfg, err := congif.InitConfig(); err != nil {
		log.Fatal("Couldn't initialize config: ", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/v1/users", handler.HandleCreateUser(apiCfg))

	log.Fatal(http.ListenAndServe(":8000", mux))
}
