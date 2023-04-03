package client

import (
	"fmt"
	"github.com/olegvelikanov/go-tcp-pow/internal/pkg/contract"
	"github.com/olegvelikanov/go-tcp-pow/internal/pkg/pow"
	"log"
	"net"
	"time"
)

const (
	connReadBufSize = 512
)

var buf = make([]byte, connReadBufSize)

func FetchQuote(addr string) ([]byte, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("connecting to remote: %s", err)
	}
	log.Printf("connected")

	challenge, err := requestChallenge(conn)
	if err != nil {
		return nil, fmt.Errorf("requesting a challenge: %s", err)
	}
	log.Printf("received challenge with difficulty=%d", challenge.CoveredBitsCount)

	t1 := time.Now()
	solution, err := challenge.Solve()
	if err != nil {
		return nil, fmt.Errorf("solving the challenge: %s", err)
	}
	log.Printf("solved challenge by %s", time.Now().Sub(t1))

	return requestService(conn, solution)
}

func requestChallenge(conn net.Conn) (*pow.Puzzle, error) {
	request := &contract.ChallengeRequest{}
	n, err := contract.Serialize(request, buf)
	if err != nil {
		return nil, fmt.Errorf("serializing challenge request: %s", err)
	}
	_, err = conn.Write(buf[:n])
	if err != nil {
		return nil, fmt.Errorf("writing to connection: %s", err)
	}

	_, err = conn.Read(buf)
	if err != nil {
		return nil, fmt.Errorf("reading from connection: %s", err)
	}
	message, err := contract.Deserialize(buf)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling challenge response: %s", err)
	}
	challengeResponse, ok := message.(*contract.ChallengeResponse)
	if !ok {
		return nil, fmt.Errorf("unexpected message type")
	}

	return challengeResponse.Puzzle, nil
}

func requestService(conn net.Conn, solution *pow.Solution) ([]byte, error) {
	request := &contract.ServiceRequest{PuzzleSolution: solution}

	n, err := contract.Serialize(request, buf)
	if err != nil {
		return nil, fmt.Errorf("serializing service request: %s", err)
	}
	_, err = conn.Write(buf[:n])
	if err != nil {
		return nil, fmt.Errorf("writing to connection: %s", err)
	}

	_, err = conn.Read(buf)
	if err != nil {
		return nil, fmt.Errorf("reading from connection: %s", err)
	}
	message, err := contract.Deserialize(buf)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling challenge response: %s", err)
	}
	serviceResponse, ok := message.(*contract.ServiceResponse)
	if !ok {
		return nil, fmt.Errorf("unexpected message type")
	}

	return serviceResponse.Quote, nil
}
