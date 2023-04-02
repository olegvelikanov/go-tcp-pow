package client

import (
	"fmt"
	"github.com/olegvelikanov/word-of-wisdom/internal/pkg/contract"
	contractpb "github.com/olegvelikanov/word-of-wisdom/internal/pkg/contract/pb"
	"github.com/olegvelikanov/word-of-wisdom/internal/pkg/pow"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
)

const (
	connReadBufSize = 2 * 1024
)

var buf = make([]byte, connReadBufSize)

func FetchQuote(port int) {
	conn, err := net.Dial("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Printf("failed to connect: %s", err)
		return
	}
	log.Printf("connected")

	challenge, err := requestChallenge(conn)
	if err != nil {
		log.Printf("Can't request challenge: %s", err)
		return
	}
	solution, err := challenge.Solve()
	if err != nil {
		log.Printf("Can't solve the challenge: %s", err)
		return
	}

	quote, err := requestService(conn, solution)
	if err != nil {
		log.Printf("Can't request quote: %s", err)
		return
	}

	log.Printf("RECEIVED QUOTE: %s", quote)
}

func requestChallenge(conn net.Conn) (*pow.Puzzle, error) {
	request := &contractpb.Message{
		Body: &contractpb.Message_ChallengeRequest{
			ChallengeRequest: &contractpb.ChallengeRequest{},
		},
	}
	bytes, err := proto.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("can't marshall challenge request: %s", err)
	}
	_, err = conn.Write(bytes)
	if err != nil {
		return nil, fmt.Errorf("can't write to connection: %s", err)
	}

	n, err := conn.Read(buf)
	if err != nil {
		return nil, fmt.Errorf("can't read from connection: %s", err)
	}
	message := &contractpb.Message{}
	err = proto.Unmarshal(buf[:n], message)
	if err != nil {
		return nil, fmt.Errorf("can't unmarshall challenge response: %s", err)
	}
	challengeResponse, ok := message.Body.(*contractpb.Message_ChallengeResponse)
	if !ok {
		return nil, fmt.Errorf("unexpected message type")
	}
	return contract.ConvPuzzleFromPb(challengeResponse.ChallengeResponse.Puzzle), nil
}

func requestService(conn net.Conn, solution *pow.Solution) (string, error) {
	request := &contractpb.Message{
		Body: &contractpb.Message_ServiceRequest{
			ServiceRequest: &contractpb.ServiceRequest{
				PuzzleSolution: contract.ConvPuzzleSolutionToPb(solution),
			},
		},
	}
	bytes, err := proto.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("can't marshall service request: %s", err)
	}
	_, err = conn.Write(bytes)
	if err != nil {
		return "", fmt.Errorf("can't write to connection: %s", err)
	}

	n, err := conn.Read(buf)
	if err != nil {
		return "", fmt.Errorf("can't read from connection: %s", err)
	}
	message := &contractpb.Message{}
	err = proto.Unmarshal(buf[:n], message)
	if err != nil {
		return "", fmt.Errorf("can't unmarshall challenge response: %s", err)
	}
	serviceResponse, ok := message.Body.(*contractpb.Message_ServiceResponse)
	if !ok {
		return "", fmt.Errorf("unexpected message type")
	}
	return serviceResponse.ServiceResponse.Quote, nil
}
