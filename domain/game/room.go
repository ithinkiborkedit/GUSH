package game

import "fmt"

type RoomID string

type Room struct {
	ID          RoomID
	Name        string
	Description string
}

func NewRoom(id RoomID, name, desc string) (*Room, error) {
	if name == "" {
		return nil, fmt.Errorf("room name cannot be empty")
	}
	return &Room{
		ID:          id,
		Name:        name,
		Description: desc,
	}, nil
}
