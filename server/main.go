package main

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("[INFO]: New connection from client: ", ws.RemoteAddr())

	s.conns[ws] = true
	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)

	for {
		n, err := ws.Read(buf)
		if err != nil {
			// Connection was closed
			if err == io.EOF {
				break
			}

			fmt.Println("[ERROR]: failed to read: ", err)
			continue
		}

		msg := buf[:n]
		s.broadcast(msg)
	}
}

func (s *Server) broadcast(msg []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			_, err := ws.Write(msg)
			if err != nil {
				fmt.Println("[ERROR]: failed to write: ", err)
			}
		}(ws)
	}
}

func main() {
	server := NewServer()
	http.Handle("/ws", websocket.Handler(server.handleWS))
	http.ListenAndServe(":3000", nil)
}
