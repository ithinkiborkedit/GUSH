package game

type PlayerMovedEvent struct {
	PlayerID PlayerID
	FromRoom RoomID
	ToRoom   RoomID
}

type PlayerSaidEvent struct {
	PlayerID PlayerID
	RoomID   RoomID
	Message  string
}
