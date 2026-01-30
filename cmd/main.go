package main

import (
	"log"
	"net/http"

	"LotteryImitationSystem/internal/handlers"
	"LotteryImitationSystem/internal/services"
	"LotteryImitationSystem/internal/storage"
)

func main() {
	userRepo := storage.NewUserRepository()
	ticketRepo := storage.NewTicketRepository()
	prizeRepo := storage.NewPrizeRepository()

	prizeChannel := make(chan string)

	service := services.NewLotteryService(
		userRepo,
		ticketRepo,
		prizeRepo,
		prizeChannel,
	)

	go service.PrizeWorker()

	userHandler := handlers.NewUserHandler(service)
	ticketHandler := handlers.NewTicketHandler(service)
	adminHandler := handlers.NewAdminHandler(service)

	mux := http.NewServeMux()
	userHandler.Register(mux)
	ticketHandler.Register(mux)
	adminHandler.Register(mux)

	fs := http.FileServer(http.Dir("./internal/frontend"))
	mux.Handle("/", fs)

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
