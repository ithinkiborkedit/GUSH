package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	serverAddr := "localhost:4000"

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	serverReader := bufio.NewReader(conn)
	userReader := bufio.NewReader(os.Stdin)
	serverWriter := bufio.NewWriter(conn)

	go readFromServer(serverReader)
	go writeToServer(userReader, serverWriter)

	select {}

}

func readFromServer(reader *bufio.Reader) {
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Disconnected from server.")
			os.Exit(0)
		}
		fmt.Print(msg)
	}
}

func writeToServer(userReader *bufio.Reader, writer *bufio.Writer) {
	for {
		input, err := userReader.ReadString('\n')
		if err != nil {
			log.Println("error reading user input: ", err)
			continue
		}
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		_, err = writer.WriteString(input + "\n")
		if err != nil {
			log.Println("error sending data to server: ", err)
			continue
		}

		writer.Flush()

		if strings.ToLower(input) == "quit" || strings.ToLower(input) == "exit" {
			fmt.Println("exiting client")
			os.Exit(0)
		}
	}

}
