package game

import "errors"

type PlayerID string

type Player struct {
	ID     PlayerID
	Name   string
	RoomID RoomID
}

func NewPlayer(id PlayerID, name string) (*Player, error) {
	if name == "" {
		return nil, errors.New("player name cannot be empty")
	}
	return &Player{
		ID:     id,
		Name:   name,
		RoomID: "",
	}, nil
}

func (p *Player) MoveToRoom(roomID RoomID) {
	p.RoomID = roomID
}
