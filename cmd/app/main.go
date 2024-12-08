package main

import (
	"Nikcase/internal/Configs"
	"Nikcase/internal/handlers"
	"Nikcase/internal/repositories"
	"Nikcase/internal/services"
	"Nikcase/pkg/database"
	"Nikcase/pkg/solutions"
	"log"
	"net/http"
)

func main() {
	solutions.ReadFileWithGoroutine()
}

func execute() {
	// Инициализация Конфигов
	configs, err := Configs.InitConfigs()
	if err != nil {
		log.Fatal(err)
		return
	}

	// 	Установка связи с Базой Данных
	db, err := database.ConnectToDB(configs)
	if err != nil {
		log.Println(err)
		return
	}

	// Цепочка Работы))
	repository := repositories.NewRepository(db)
	service := services.NewService(repository)
	handler := handlers.NewHandlers(service)
	log.Println("Main Work is done!")

	// Инициализация Маршрутизатора и передачча поинта в Сервер
	mux := InitRouters(handler)
	log.Println("Init MUX is done successfully!")

	srv := http.Server{
		Addr:    configs.Server.Host + configs.Server.Port,
		Handler: mux,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Println("Listen and Serv ERROR!!!\n", err)
		return
	}

}
