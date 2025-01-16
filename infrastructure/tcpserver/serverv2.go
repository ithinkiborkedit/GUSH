package tcpserver

import (
	"fmt"
	"log"
	"net"

	"github.com/ithinkiborkedit/GUSH.git"
	app "github.com/ithinkiborkedit/GUSH.git/application/game"
	domain "github.com/ithinkiborkedit/GUSH.git/domain/game"
)

type TCPServer struct {
	GUSHUseCase *app.GUSHUseCase
}

func (s *TCPServer) Listen() {
	port := ":4000"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen on port %s", port)
	}
	defer listener.Close()

	log.Printf("GUSH Server listening on port %s", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Panicln("accept error: ", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *TCPServer) handleConnection(conn net.Conn) {
	defer conn.Close()

	playerID := domain.PlayerID(conn.RemoteAddr().String())
	player, _ := domain.NewPlayer(playerID, "Guest-"+string(playerID))
	_ = s.GUSHUseCase.PlayerRepo.Save(player)

	player.MoveToRoom("room-1")
	_ = s.GUSHUseCase.PlayerRepo.Save(player)

	welcomeMsg := &GUSH.ServerMessage{
		Payload: &GUSH.ServerMessage_SystemMsg{
			SystemMsg: &GUSH.SystemMsg{
				Text: fmt.Sprintf("Welcome, %s! You are in %s.", player.Name, player.RoomID),
			},
		},
	}

	_ = WriteProto(conn, welcomeMsg)

	for {
		cmd := &GUSH.Command{}
		err := ReadProto(conn, cmd)

		if err != nil {
			log.Printf("Connection from %s closed or error: %v\n", conn.RemoteAddr(), err)
			return
		}

		switch cmd.Type {
		case "move":
			moveErr := s.GUSHUseCase.HandleMove(domain.MoveCommand{
				PlayerID: player.ID,
				RoomID:   domain.RoomID(cmd.Payload),
			})
			if moveErr != nil {
				resp := &GUSH.ServerMessage{
					Payload: &GUSH.ServerMessage_SystemMsg{
						SystemMsg: &GUSH.SystemMsg{
							Text: fmt.Sprintf("move failed: %v", moveErr),
						},
					},
				}
				_ = WriteProto(conn, resp)
				continue
			}

			newRoom, _ := s.GUSHUseCase.RoomRepo.FindByID(domain.RoomID(cmd.Payload))
			if newRoom != nil {
				roomMsg := &GUSH.ServerMessage{
					Payload: &GUSH.ServerMessage_RoomUpdate{
						RoomUpdate: &GUSH.RoomUpdate{
							RoomId:      string(newRoom.ID),
							RoomName:    newRoom.Name,
							Description: newRoom.Description,
						},
					},
				}
				_ = WriteProto(conn, roomMsg)
			}
		case "say":
			err := s.GUSHUseCase.HandleSay(domain.SayCommand{
				PlayerID: player.ID,
				Message:  cmd.Payload,
			})
			if err != nil {
				resp := &GUSH.ServerMessage{
					Payload: &GUSH.ServerMessage_SystemMsg{
						SystemMsg: &GUSH.SystemMsg{
							Text: fmt.Sprintf("say failed: %v", err),
						},
					},
				}
				_ = WriteProto(conn, resp)
				continue
			}
			echo := &GUSH.ServerMessage{
				Payload: &GUSH.ServerMessage_Chat{
					Chat: &GUSH.ChatMessage{
						PlayerName: player.Name,
						Text:       cmd.Payload,
					},
				},
			}
			_ = WriteProto(conn, echo)
		default:
			unknown := &GUSH.ServerMessage{
				Payload: &GUSH.ServerMessage_SystemMsg{
					SystemMsg: &GUSH.SystemMsg{
						Text: fmt.Sprintf("unknown command type %s", cmd.Type),
					},
				},
			}
			_ = WriteProto(conn, unknown)
		}
	}
}
