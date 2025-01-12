package models

import (
	"fmt"
	"sync"
)

type Room struct {
	Name        string
	Description string
	Users       map[string]*User
	Exits       map[string]*Room
	lock        sync.RWMutex
}

func NewRoom(name, description string) *Room {
	return &Room{
		Name:        name,
		Description: description,
		Users:       make(map[string]*User),
	}
}

func (r *Room) AddUser(u *User) {
	r.lock.Lock()
	defer r.lock.Unlock()
	if r.Users == nil {
		r.Users = make(map[string]*User)
	}
	r.Users[u.Username] = u

	r.Broadcast(fmt.Sprintf("%s has entered the room", u.Username), u)
}

func (r *Room) RemoveUser(u *User) {
	r.lock.Lock()
	defer r.lock.Unlock()
	delete(r.Users, u.Username)

	r.Broadcast(fmt.Sprintf("%s has left the room", u.Username), u)
}

func (r *Room) Broadcast(message string, sender *User) {
	fmt.Println(message)
}

func (r *Room) SetExit(direction string, room *Room) {
	if r.Exits == nil {
		r.Exits = make(map[string]*Room)
	}
	r.Exits[direction] = room
}
