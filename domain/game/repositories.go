package game

type PlayerRepository interface {
	FindByID(id PlayerID) (*Player, error)
	FindAllInRoom(roomID RoomID) ([]*Player, error)
	Save(player *Player) error
}

type RoomRepository interface {
	FindByID(id RoomID) (*Room, error)
	Save(room *Room) error
}
