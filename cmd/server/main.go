package main

import (
	"github.com/ithinkiborkedit/GUSH.git/pkg/commands"
	"github.com/ithinkiborkedit/GUSH.git/pkg/game"
	"github.com/ithinkiborkedit/GUSH.git/pkg/network"
	"github.com/ithinkiborkedit/GUSH.git/pkg/storage"
	"github.com/ithinkiborkedit/GUSH.git/pkg/user"
)

func main() {
	storage.InitDB("GUSH.db")
	defer storage.CloseDB()

	world := game.Initializeworld()

	userManager := user.NewUserManager()

	dispatcher := commands.NewDispatcher()

	commands.RegisterCommands(dispatcher)

	server := network.NewServer(":4000", dispatcher, world, userManager)

	server.Start()
}
