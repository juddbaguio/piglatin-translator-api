package main

import (
	"log"
	"piglatin-translator-api/infrastructure/db"
	"piglatin-translator-api/repo"
	"piglatin-translator-api/transport/http"
	"piglatin-translator-api/usecase"
)

func main() {
	postgresDB, err := db.SetupDatabase()
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer postgresDB.Close()

	piglatinRepo := repo.NewPiglatinRepo(postgresDB)
	piglatinUC := usecase.NewPiglatinUsecase(piglatinRepo)
	server := http.NewServer(piglatinUC)

	log.Println("server is now running at: ", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Println("Error starting server: ", err.Error())
		return
	}
}
