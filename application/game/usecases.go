package game

import (
	"fmt"

	domain "github.com/ithinkiborkedit/GUSH.git/domain/game"
)

type GUSHUseCase struct {
	PlayerRepo domain.PlayerRepository
	RoomRepo   domain.RoomRepository
	Broadcast  *domain.BroadcastService
}

func NewGUSHUseCase(pRepo domain.PlayerRepository, rRepo domain.RoomRepository, bCast *domain.BroadcastService) *GUSHUseCase {
	return &GUSHUseCase{
		PlayerRepo: pRepo,
		RoomRepo:   rRepo,
		Broadcast:  bCast,
	}
}

func (uc *GUSHUseCase) HandleSay(cmd domain.SayCommand) error {
	player, err := uc.PlayerRepo.FindByID(cmd.PlayerID)
	if err != nil {
		return fmt.Errorf("player not found: %v", err)
	}

	message := fmt.Sprintf("%s says: %s", player.Name, cmd.Message)
	return uc.Broadcast.BroadcastToRoom(player.RoomID, message)
}

func (uc *GUSHUseCase) HandleMove(cmd domain.MoveCommand) error {
	player, err := uc.PlayerRepo.FindByID(cmd.PlayerID)
	if err != nil {
		return fmt.Errorf("player not found: %v", err)
	}

	fromRoom := player.RoomID

	_, err = uc.RoomRepo.FindByID(cmd.RoomID)
	if err != nil {
		return fmt.Errorf("room not found")
	}

	player.MoveToRoom(cmd.RoomID)
	if err := uc.PlayerRepo.Save(player); err != nil {
		return fmt.Errorf("unable to save player state %v", err)
	}

	leaveMsg := fmt.Sprintf("%s leaves for %s", player.Name, cmd.RoomID)
	if err := uc.Broadcast.BroadcastToRoom(fromRoom, leaveMsg); err != nil {
		return fmt.Errorf("[ERROR %v]: ", err)
	}
	arriveMsg := fmt.Sprintf("%s arrives from %s", player.Name, fromRoom)
	if err := uc.Broadcast.BroadcastToRoom(cmd.RoomID, arriveMsg); err != nil {
		return fmt.Errorf("[ERROR %v]: ", err)
	}

	// moveMsg := fmt.Sprintf("%s moves from %s to %s", player.Name, fromRoom, cmd.RoomID)
	// return uc.Broadcast.BroadcastToRoom(fromRoom, moveMsg)
	return nil
}
