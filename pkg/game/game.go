package game

import (
	"log"

	"github.com/ithinkiborkedit/GUSH.git/pkg/models"
)

type World struct {
	Rooms map[string]*models.Room
}

func NewWorld() *World {
	return &World{
		Rooms: make(map[string]*models.Room),
	}
}

func (w *World) AddRoom(room *models.Room) {
	if w.Rooms == nil {
		w.Rooms = make(map[string]*models.Room)
	}
	w.Rooms[room.Name] = room
}

func (w *World) GetRoom(name string) *models.Room {
	return w.Rooms[name]
}

func Initializeworld() *World {
	world := NewWorld()

	startRoom := models.NewRoom("start", "you are in the starting room Exits are north and east.")
	secondRoom := models.NewRoom("second", "this is the second room, Exits are south")
	eastRoom := models.NewRoom("east", "You've moved east. there's nothing here")

	startRoom.SetExit("north", secondRoom)
	startRoom.SetExit("east", eastRoom)
	secondRoom.SetExit("south", startRoom)

	world.AddRoom(startRoom)
	world.AddRoom(secondRoom)
	world.AddRoom(eastRoom)

	log.Println("World initialized with rooms: start, second, east")

	return world
}
