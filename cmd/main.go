package main

import (
	"log"

	app "github.com/ithinkiborkedit/GUSH.git/application/game"
	domain "github.com/ithinkiborkedit/GUSH.git/domain/game"
	"github.com/ithinkiborkedit/GUSH.git/infrastructure/storage"
	"github.com/ithinkiborkedit/GUSH.git/infrastructure/tcpserver"
)

func main() {

	playerRepo := storage.NewInMemoryPlayerRepo()
	roomRepo := storage.NewInMemoryRoomRepo()

	lobby, _ := domain.NewRoom("room-0", "lobby", "You see tons of people getting ready to explore!")
	startRoom, _ := domain.NewRoom("room-1", "Starting Room", "A small, cozy room.")
	endRoom, _ := domain.NewRoom("room-2", "Your Final resting place.", "You fell into a firey pit and died")

	log.Printf("Initializing room(s): %s", startRoom.ID)
	roomRepo.Save(startRoom)
	roomRepo.Save(endRoom)
	roomRepo.Save(lobby)

	broadcastService := &domain.BroadcastService{
		PlayerRepo: playerRepo,
		RoomRepo:   roomRepo,
	}

	gushUseCase := app.NewGUSHUseCase(playerRepo, roomRepo, broadcastService)

	server := &tcpserver.TCPServer{GUSHUseCase: gushUseCase}
	log.Println("Starting GUSH server...")

	server.Listen()

}
