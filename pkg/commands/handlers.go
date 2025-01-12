package commands

import (
	"fmt"

	"strings"

	"github.com/ithinkiborkedit/GUSH.git/pkg/game"
	"github.com/ithinkiborkedit/GUSH.git/pkg/models"
)

func RegisterCommands(d *Dispatcher) {
	d.Register("look", LookCommand)
	d.Register("say", SayCommand)
	d.Register("go", GoCommand)
	d.Register("quit", QuitCommand)
}

func LookCommand(args []string, usr *models.User, world *game.World) string {
	room := usr.CurrentRoom
	desc := room.Description + "\n Exits: " + strings.Join(getExitNames(room), ", ")
	return desc
}

func getExitNames(room *models.Room) []string {
	exits := []string{}
	for direction := range room.Exits {
		exits = append(exits, direction)
	}
	return exits
}

func SayCommand(args []string, usr *models.User, world *game.World) string {
	if len(args) == 0 {
		return "Say what?"
	}

	message := strings.Join(args, " ")
	room := usr.CurrentRoom

	for _, u := range room.Users {
		if u.Username != usr.Username {
			fmt.Printf("%s says: %s\n", usr.Username, message)
		}
	}

	return fmt.Sprintf("You say: %s", message)
}

func GoCommand(args []string, usr *models.User, world *game.World) string {
	if len(args) == 0 {
		return "Go Where?"
	}

	direction := strings.ToLower(args[0])
	currentRoom := usr.CurrentRoom

	nextRoom, exists := currentRoom.Exits[direction]
	if !exists {
		return "You cannot go that way"
	}

	currentRoom.RemoveUser(usr)
	nextRoom.AddUser(usr)
	usr.CurrentRoom = nextRoom

	return nextRoom.Description
}

func QuitCommand(args []string, usr *models.User, world *game.World) string {
	return "Good bye!"
}
