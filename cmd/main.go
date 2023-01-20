package main

import (
	"log"
	"piglatin-translator-api/infrastructure/db"
	"piglatin-translator-api/repo"
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

	log.Println(piglatinUC.Translate("we are on a riding,in,tandem"))
	log.Println(piglatinUC.GetTranslationRequests(1))
}
