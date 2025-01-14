package game

import "fmt"

type BroadcastService struct {
	PlayerRepo PlayerRepository
	RoomRepo   RoomRepository
}

func (b *BroadcastService) BroadcastToRoom(roomID RoomID, message string) error {
	players, err := b.PlayerRepo.FindAllInRoom(roomID)
	if err != nil {
		return err
	}

	for _, p := range players {
		fmt.Printf("Sending to %s: %s\n", p.Name, message)
	}

	return nil
}
