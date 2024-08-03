package main

import (
	"log"
	"net/http"
	
	"github.com/balagrivine/go_auth/config"
	"github.com/balagrivine/go_auth/handler"
	"github.com/joho/godotenv"
)

func main() {

	// Load environmental variables
	godotenv.Load(".env")

	// Initialize the ocnfiguration
	apiCfg, err := config.InitConfig()
	if err != nil {
		log.Fatal("Couldn't initialize config: ", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/v1/users", handler.HandleCreateUser(apiCfg))

	log.Fatal(http.ListenAndServe(":8080", mux))
}
