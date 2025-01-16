package tcpserver

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/ithinkiborkedit/GUSH.git/application/game"
	domain "github.com/ithinkiborkedit/GUSH.git/domain/game"
)

type TCPServerv1 struct {
	GUSHUseCase *game.GUSHUseCase
}

func (s *TCPServerv1) Listen(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen on port %s", port)
	}

	defer listener.Close()

	log.Printf("Listening on port %s\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *TCPServerv1) handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	playerID := domain.PlayerID(conn.RemoteAddr().String())
	guestString := fmt.Sprintf("Guest-%s", playerID)
	tempPlayer, _ := domain.NewPlayer(playerID, guestString)

	tempPlayer.RoomID = "room-0"

	s.GUSHUseCase.PlayerRepo.Save(tempPlayer)

	fmt.Fprintf(conn, "Welcom to GUSH!\n")

	for {
		raw, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Connection lost: %v", err)
			return
		}

		cmdLine := strings.TrimSpace(raw)
		if cmdLine == "" {
			continue
		}

		fields := strings.SplitN(cmdLine, " ", 2)
		cmd := fields[0]
		args := ""
		if len(fields) > 1 {
			args = fields[1]
		}

		switch cmd {
		case "/say":
			s.GUSHUseCase.HandleSay(domain.SayCommand{
				PlayerID: playerID,
				Message:  args,
			})
		case "/move":
			err := s.GUSHUseCase.HandleMove(domain.MoveCommand{
				PlayerID: playerID,
				RoomID:   domain.RoomID(args),
			})
			if err != nil {
				fmt.Fprintf(conn, "Failed to move: %v\n", err)
				return
			}
			fmt.Fprintf(conn, "You moved to %s. \n", args)
		default:
			fmt.Fprintf(conn, "unknown command: %s\n", cmd)
		}
	}
}
