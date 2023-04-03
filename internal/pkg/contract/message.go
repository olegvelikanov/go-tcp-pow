package contract

import (
	"fmt"
)

var (
	ErrNotEnoughBytes = fmt.Errorf("not enough bytes in slice")
	ErrUnsupportedMsg = fmt.Errorf("unsupported message type")
)

type Message interface {
	isMessage()
}

func Serialize(message Message, b []byte) (int, error) {
	if len(b) < 1 {
		return 0, ErrNotEnoughBytes
	}
	switch m := message.(type) {
	case *ChallengeRequest:
		b[0] = 0x01
		n, err := serializeChallengeRequest(m, b[1:])
		return n + 1, err
	case *ChallengeResponse:
		b[0] = 0x02
		n, err := serializeChallengeResponse(m, b[1:])
		return n + 1, err
	case *ServiceRequest:
		b[0] = 0x03
		n, err := serializeServiceRequest(m, b[1:])
		return n + 1, err
	case *ServiceResponse:
		b[0] = 0x04
		n, err := serializeServiceResponse(m, b[1:])
		return n + 1, err
	default:
		return 0, ErrUnsupportedMsg
	}
}

func Deserialize(b []byte) (Message, error) {
	if len(b) < 1 {
		return nil, ErrNotEnoughBytes
	}
	switch b[0] {
	case 0x01:
		return deserializeChallengeRequest(b[1:])
	case 0x02:
		return deserializeChallengeResponse(b[1:])
	case 0x03:
		return deserializeServiceRequest(b[1:])
	case 0x04:
		return deserializeServiceResponse(b[1:])
	default:
		return nil, ErrUnsupportedMsg
	}
}
