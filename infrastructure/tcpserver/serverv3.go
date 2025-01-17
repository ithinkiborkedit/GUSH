package tcpserver

import (
	"net"
	"time"
)

type Server struct {
	address      string
	port         int
	readTimeOut  time.Duration
	writeTimeOut time.Duration
	protocol     string
	listener     net.Listener
}

type IServer interface {
	Start()
	Stop()
	HotReload()
}

type Option func(*Server)

func withAddress(addr string) Option {
	return func(s *Server) {
		s.address = addr
	}
}

func withPort(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}

func withReadTimeOut(rt time.Duration) Option {
	return func(s *Server) {
		s.readTimeOut = rt
	}
}

func withWriteTimeOut(wt time.Duration) Option {
	return func(s *Server) {
		s.readTimeOut = wt
	}
}

func withProtocol(protoc string) Option {
	return func(s *Server) {
		s.protocol = protoc
	}
}
