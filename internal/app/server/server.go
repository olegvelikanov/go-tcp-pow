package server

import (
	"errors"
	"fmt"
	"github.com/olegvelikanov/go-tcp-pow/internal/pkg/contract"
	"io"
	"log"
	"net"
)

const (
	connReadBufSize = 512
)

type Server struct {
	listener net.Listener
	app      Application
}

func StartServer(config *Config) (*Server, error) {
	listener, err := net.Listen("tcp", config.Address)
	if err != nil {
		return nil, err
	}
	app, err := NewApp(config)
	if err != nil {
		return nil, fmt.Errorf("creating app: %s", err)
	}
	s := &Server{
		listener: listener,
		app:      app,
	}
	go s.serve()
	log.Printf("started tcp server serving at %s", config.Address)
	return s, nil
}

func (s *Server) serve() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return
			}
			log.Printf("error accepting connection: %s", err)
			continue
		}
		go s.serveConnection(conn)
	}
}

func (s *Server) serveConnection(conn net.Conn) {
	defer func() {
		conn.Close()
		log.Printf("connection closed: %s", conn.RemoteAddr())
	}()
	log.Printf("connection established: %s", conn.RemoteAddr())

	bufPtr := s.getBuf()
	defer s.freeBuf(bufPtr)
	buf := *bufPtr

	for {
		_, err := conn.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Printf("client sent EOF")
			} else {
				log.Printf("error reading message: %s", err)
			}
			return
		}
		err = s.handleMessage(conn, buf)
		if err != nil {
			log.Printf("error handling message: %s", err)
			return
		}
	}
}

func (s *Server) handleMessage(conn net.Conn, buf []byte) error {
	message, err := contract.Deserialize(buf)
	if err != nil {
		return fmt.Errorf("deserializing message: %s", err)
	}

	switch msg := message.(type) {
	case *contract.ChallengeRequest:
		return s.onChallengeRequest(conn, buf, msg)
	case *contract.ServiceRequest:
		return s.onServiceRequest(conn, buf, msg)
	default:
		return fmt.Errorf("unexpected message received")
	}
}

func (s *Server) onChallengeRequest(conn net.Conn, buf []byte, _ *contract.ChallengeRequest) error {
	log.Printf("challenge requested")
	puzzle := s.app.onChallengeRequest()
	return s.sendResponse(conn, buf, &contract.ChallengeResponse{Puzzle: puzzle})
}

func (s *Server) onServiceRequest(conn net.Conn, buf []byte, request *contract.ServiceRequest) error {
	log.Printf("service requested")
	quote, err := s.app.onServiceRequest(request.PuzzleSolution)
	if err != nil {
		return fmt.Errorf("requesting service: %s", err)
	}
	return s.sendResponse(conn, buf, &contract.ServiceResponse{Quote: quote})
}

func (s *Server) sendResponse(conn net.Conn, buf []byte, message contract.Message) error {
	n, err := contract.Serialize(
		message,
		buf,
	)
	if err != nil {
		return fmt.Errorf("serializing message: %s", err)
	}
	_, err = conn.Write(buf[:n])
	if err != nil {
		return fmt.Errorf("writing message to connection: %s", err)
	}
	return nil
}

func (s *Server) Stop() {
	log.Printf("stopping tcp server")
	err := s.listener.Close()
	if err != nil {
		log.Printf("can't stop tcp server: %s", err)
		return
	}
	log.Printf("server is stopped")
}

// TODO: use sync.Pool here

func (s *Server) getBuf() *[]byte {
	bytes := make([]byte, connReadBufSize)
	return &bytes
}
func (s *Server) freeBuf(ptr *[]byte) {

}
