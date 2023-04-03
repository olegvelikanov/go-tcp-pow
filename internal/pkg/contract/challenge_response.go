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
// |        16 | covered h1 hash    | 32 bytes
// |        48 | covered bits count | 1 byte
// |        49 | h2 hash            | 32 bytes
// |        81 |                    |

const (
	chRspSecondsOffset          = 0
	chRspNanosOffset            = 8
	chRspHashOneOffset          = 16
	chRspCoveredBitsCountOffset = 48
	chRspHashTwoOffset          = 49
	chRspTotalSize              = 81
)

type ChallengeResponse struct {
	Puzzle *pow.Puzzle
}

func (*ChallengeResponse) isMessage() {}

func serializeChallengeResponse(m *ChallengeResponse, b []byte) (int, error) {
	if len(b) < chRspTotalSize {
		return 0, ErrNotEnoughBytes
	}

	seconds := m.Puzzle.Timestamp.Unix()
	binary.LittleEndian.PutUint64(b[chRspSecondsOffset:chRspNanosOffset], uint64(seconds))

	nanos := m.Puzzle.Timestamp.Nanosecond()
	binary.LittleEndian.PutUint64(b[chRspNanosOffset:chRspHashOneOffset], uint64(nanos))

	copy(b[chRspHashOneOffset:chRspCoveredBitsCountOffset], m.Puzzle.CoveredHashOne)

	b[chRspCoveredBitsCountOffset] = m.Puzzle.CoveredBitsCount

	copy(b[chRspHashTwoOffset:chRspTotalSize], m.Puzzle.HashTwo)

	return chRspTotalSize, nil
}

func deserializeChallengeResponse(b []byte) (*ChallengeResponse, error) {
	if len(b) < chRspTotalSize {
		return nil, ErrNotEnoughBytes
	}
	seconds := binary.LittleEndian.Uint64(b[chRspSecondsOffset:chRspNanosOffset])
	nanos := binary.LittleEndian.Uint64(b[chRspNanosOffset:chRspHashOneOffset])

	return &ChallengeResponse{
		Puzzle: &pow.Puzzle{
			Timestamp:        time.Unix(int64(seconds), int64(nanos)).UTC(),
			CoveredHashOne:   b[chRspHashOneOffset:chRspCoveredBitsCountOffset],
			CoveredBitsCount: b[chRspCoveredBitsCountOffset],
			HashTwo:          b[chRspHashTwoOffset:chRspTotalSize],
		},
	}, nil
}
