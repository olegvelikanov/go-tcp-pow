package server

import (
	"errors"
	"fmt"
	"github.com/olegvelikanov/word-of-wisdom/internal/pkg/contract"
	contractpb "github.com/olegvelikanov/word-of-wisdom/internal/pkg/contract/pb"
	"log"
	"net"
)

const (
	connReadBufSize = 2 * 1024
)

type Server struct {
	port     int
	listener net.Listener
	app      Application
}

func StartServer(port int) (*Server, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}
	s := &Server{
		port:     port,
		listener: listener,
		app:      NewWordOfWisdomApp(),
	}
	go s.serve()
	log.Printf("Started tcp server serving at %d", s.port)
	return s, nil
}

func (s *Server) serve() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return
			}
			log.Printf("Error accepting connection: %s", err)
			continue
		}
		go s.serveConnection(conn)
	}
}

func (s *Server) serveConnection(conn net.Conn) {
	defer func() {
		conn.Close()
		log.Printf("Connection closed: %s", conn.RemoteAddr())
	}()
	log.Printf("Connection established: %s", conn.RemoteAddr())

	buf := make([]byte, connReadBufSize)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			return
		}
		if err != nil {
			log.Printf("Error reading from connection(%s): %s", conn.RemoteAddr(), err)
			break
		}
		s.handleMessage(conn, buf[:n])
	}
}

func (s *Server) handleMessage(conn net.Conn, m []byte) {
	message, err := contract.Deserialize(m)
	if err != nil {
		log.Printf("Can't deserialize message: %s", err)
		return
	}

	switch message.Body.(type) {
	case *contractpb.Message_ChallengeRequest:
		s.onChallengeRequest(conn, message.Body.(*contractpb.Message_ChallengeRequest))
	case *contractpb.Message_ServiceRequest:
		s.onServiceRequest(conn, message.Body.(*contractpb.Message_ServiceRequest))
	default:

	}
}

func (s *Server) onChallengeRequest(conn net.Conn, _ *contractpb.Message_ChallengeRequest) {
	puzzle := s.app.onChallengeRequest()
	bytes, err := contract.Serialize(&contractpb.Message{
		Body: &contractpb.Message_ChallengeResponse{
			ChallengeResponse: &contractpb.ChallengeResponse{
				Puzzle: contract.ConvPuzzleToPb(puzzle),
			},
		},
	})
	if err != nil {
		log.Printf("Can't serialize message: %s", err)
		return
	}
	_, err = conn.Write(bytes)
	if err != nil {
		log.Printf("Can't write message to connection: %s", err)
		return
	}
}

func (s *Server) onServiceRequest(conn net.Conn, request *contractpb.Message_ServiceRequest) {
	quote, err := s.app.onServiceRequest(contract.ConvPuzzleSolutionFromPb(request.ServiceRequest.PuzzleSolution))
	if err != nil {
		log.Printf("Error processing service request: %s", err)
	}
	bytes, err := contract.Serialize(&contractpb.Message{
		Body: &contractpb.Message_ServiceResponse{
			ServiceResponse: &contractpb.ServiceResponse{
				Quote: quote,
			},
		},
	})
	if err != nil {
		log.Printf("Can't serialize message: %s", err)
		return
	}
	_, err = conn.Write(bytes)
	if err != nil {
		log.Printf("Can't write message to connection: %s", err)
		return
	}
}

func (s *Server) Stop() {
	log.Printf("Stopping tcp server")
	err := s.listener.Close()
	if err != nil {
		log.Printf("Can't stop tcp server: %s", err)
		return
	}
	log.Printf("Server is stopped")
}
