package main

import (
	"log"
	"net/http"

	"github.com/nunusavi/task-manager/internal/respository"
	"github.com/nunusavi/task-manager/internal/config"
	"github.com/nunusavi/task-manager/internal/handler"
)

func main() {

	cfg := config.LoadConfig()
	respository.InitDB(cfg)
	
	r := handler.NewRouter()
	
	log.Println(
		"server listening on http://localhost:8080",
	)
	log.Fatal(http.ListenAndServe(":8080", r))

}
