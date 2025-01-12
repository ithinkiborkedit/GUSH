package network

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/ithinkiborkedit/GUSH.git/pkg/commands"
	"github.com/ithinkiborkedit/GUSH.git/pkg/game"
	"github.com/ithinkiborkedit/GUSH.git/pkg/user"
)

type Server struct {
	Address     string
	Dispatcher  *commands.Dispatcher
	World       *game.World
	UserManager *user.UserManager
}

func NewServer(address string, dispatcher *commands.Dispatcher, world *game.World, userManager *user.UserManager) *Server {
	return &Server{
		Address:     address,
		Dispatcher:  dispatcher,
		World:       world,
		UserManager: userManager,
	}
}

func (s *Server) Start() {
	listner, err := net.Listen("tcp", s.Address)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer listner.Close()
	log.Printf("GUSH server listening on %s\n", s.Address)

	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %v", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	writer.WriteString("Welcome to GUSH!\n")
	writer.WriteString("Please log in \nUsername: ")

	writer.Flush()

	username, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Error reading username: %v", err)
		return
	}
	username = strings.TrimSpace(username)

	usr, err := s.UserManager.Authenticate(username)
	if err != nil {
		writer.WriteString("Authentication failed.\n")
		writer.Flush()
		return
	}

	startingRoom := s.World.GetRoom("start")
	if startingRoom == nil {
		writer.WriteString("Starting room not found!\n")
		writer.Flush()
		return
	}

	usr.CurrentRoom = startingRoom
	startingRoom.AddUser(usr)

	writer.WriteString(fmt.Sprintf("Welcome! %s", usr.Username))
	writer.WriteString(usr.CurrentRoom.Description + "\n")

	writer.Flush()

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("connection closed for user %s", usr.Username)
			usr.CurrentRoom.RemoveUser(usr)
			break
		}

		input = strings.TrimSpace(input)
		if input == "" {
			writer.WriteString("> ")
			writer.Flush()
			continue
		}

		response := s.Dispatcher.Dispatch(input, usr, s.World)
		writer.WriteString(response + "\n> ")
		writer.Flush()

		if strings.ToLower(input) == "quit" {
			usr.CurrentRoom.RemoveUser(usr)
			break
		}
	}
}
