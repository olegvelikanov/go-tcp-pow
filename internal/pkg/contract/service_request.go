package contract

import (
	"encoding/binary"
	"github.com/olegvelikanov/go-tcp-pow/internal/pkg/pow"
	"time"
)

// | offset    | description        | length
// |-----------|--------------------|-----------
// |         0 | timestamp seconds  | 8 bytes
// |         8 | timestamp nanos    | 8 bytes
// |        16 | h1 hash            | 32 bytes
// |        48 |                    |

const (
	srvRqSecondsOffset = 0
	srvRqNanosOffset   = 8
	srvRqHashOneOffset = 16
	srvRqTotalSize     = 48
)

type ServiceRequest struct {
	PuzzleSolution *pow.Solution
}

func (*ServiceRequest) isMessage() {}

func serializeServiceRequest(m *ServiceRequest, b []byte) (int, error) {
	if len(b) < srvRqTotalSize {
		return 0, ErrNotEnoughBytes
	}

	seconds := m.PuzzleSolution.Timestamp.Unix()
	binary.LittleEndian.PutUint64(b[srvRqSecondsOffset:srvRqNanosOffset], uint64(seconds))

	nanos := m.PuzzleSolution.Timestamp.Nanosecond()
	binary.LittleEndian.PutUint64(b[srvRqNanosOffset:srvRqHashOneOffset], uint64(nanos))

	copy(b[srvRqHashOneOffset:srvRqTotalSize], m.PuzzleSolution.HashOne)

	return srvRqTotalSize, nil
}

func deserializeServiceRequest(b []byte) (*ServiceRequest, error) {
	if len(b) < srvRqTotalSize {
		return nil, ErrNotEnoughBytes
	}
	seconds := binary.LittleEndian.Uint64(b[srvRqSecondsOffset:srvRqNanosOffset])
	nanos := binary.LittleEndian.Uint64(b[srvRqNanosOffset:srvRqHashOneOffset])

	return &ServiceRequest{
		PuzzleSolution: &pow.Solution{
			Timestamp: time.Unix(int64(seconds), int64(nanos)).UTC(),
			HashOne:   b[srvRqHashOneOffset:srvRqTotalSize],
		},
	}, nil
}
