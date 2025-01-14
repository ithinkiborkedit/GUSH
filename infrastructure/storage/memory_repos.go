package storage

import (
	"errors"
	"sync"

	domain "github.com/ithinkiborkedit/GUSH.git/domain/game"
)

type InMemoryPlayerRepo struct {
	mu      sync.RWMutex
	players map[domain.PlayerID]*domain.Player
}

func NewInMemoryPlayerRepo() *InMemoryPlayerRepo {
	return &InMemoryPlayerRepo{
		players: make(map[domain.PlayerID]*domain.Player),
	}
}

func (r *InMemoryPlayerRepo) FindByID(id domain.PlayerID) (*domain.Player, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	p, ok := r.players[id]
	if !ok {
		return nil, errors.New("player not found")
	}

	return p, nil
}

func (r *InMemoryPlayerRepo) FindAllInRoom(roomID domain.RoomID) ([]*domain.Player, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*domain.Player

	for _, p := range r.players {
		if p.RoomID == roomID {
			result = append(result, p)
		}
	}
	return result, nil
}

func (r *InMemoryPlayerRepo) Save(player *domain.Player) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.players[player.ID] = player

	return nil
}

//____________

type InMemoryRoomRepo struct {
	mu    sync.RWMutex
	rooms map[domain.RoomID]*domain.Room
}

func NewInMemoryRoomRepo() *InMemoryRoomRepo {
	return &InMemoryRoomRepo{
		rooms: make(map[domain.RoomID]*domain.Room),
	}
}

func (r *InMemoryRoomRepo) FindByID(id domain.RoomID) (*domain.Room, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	room, ok := r.rooms[id]
	if !ok {
		return nil, errors.New("player not found")
	}

	return room, nil
}

func (r *InMemoryRoomRepo) Save(room *domain.Room) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.rooms[room.ID] = room

	return nil
}
