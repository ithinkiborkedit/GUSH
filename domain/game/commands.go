package game

type SayCommand struct {
	PlayerID PlayerID
	Message  string
}

type MoveCommand struct {
	PlayerID PlayerID
	RoomID   RoomID
}
