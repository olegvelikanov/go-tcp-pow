package pow

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"github.com/benbjohnson/clock"
	"hash"
	"time"
)

var ErrSolutionNotFound = fmt.Errorf("puzzle solution not found")

type Puzzle struct {
	Timestamp        time.Time
	CoveredHashOne   []byte
	CoveredBitsCount uint8
	HashTwo          []byte
}

type Solution struct {
	Timestamp time.Time
	HashOne   []byte
}

func NewPuzzle(bitsCount uint8, secret []byte, clock clock.Clock) *Puzzle {
	timestamp := clock.Now()
	hashOne := calculateHashOne(timestamp, secret)
	hashTwo := calculateHash(hashOne)

	return &Puzzle{
		Timestamp:        timestamp,
		CoveredHashOne:   coverTrailingBits(hashOne, int(bitsCount)),
		CoveredBitsCount: bitsCount,
		HashTwo:          hashTwo,
	}
}

func calculateHashOne(timestamp time.Time, secret []byte) []byte {
	ts := make([]byte, 16)
	binary.LittleEndian.PutUint64(ts[0:8], uint64(timestamp.Unix()))
	binary.LittleEndian.PutUint64(ts[8:16], uint64(timestamp.Nanosecond()))
	return calculateHash(ts, secret)
}

func calculateHash(args ...[]byte) []byte {
	h := getHash()
	for _, arg := range args {
		h.Write(arg)
	}
	return h.Sum(nil)
}

func coverTrailingBits(bytes []byte, bitsToReset int) []byte {
	idx := len(bytes) - 1 - (bitsToReset / 8)
	for i := range bytes {
		if i > idx {
			bytes[i] = 0
		} else if i == idx {
			bytes[i] = (bytes[i] >> (bitsToReset % 8)) << (bitsToReset % 8)
		}
	}
	return bytes
}

func (p *Puzzle) Solve() (*Solution, error) {
	hashOne := make([]byte, len(p.CoveredHashOne))
	copy(hashOne, p.CoveredHashOne)
	found := false

	n := 2 << p.CoveredBitsCount
	for i := 0; i < n; i++ {
		if bytes.Equal(p.HashTwo, calculateHash(hashOne)) {
			found = true
			break
		}
		increment(hashOne)
	}
	if !found {
		return nil, ErrSolutionNotFound
	}

	return &Solution{
		Timestamp: p.Timestamp,
		HashOne:   hashOne,
	}, nil
}

func increment(bytes []byte) []byte {
	i := len(bytes) - 1
	for bytes[i] == byte(0xff) && i > 0 {
		bytes[i] = byte(0x00)
		i--
	}
	if i >= 0 {
		bytes[i] = bytes[i] + 1
	}
	return bytes
}

func (s *Solution) IsValid(secret []byte) bool {
	if !bytes.Equal(s.HashOne, calculateHashOne(s.Timestamp, secret)) {
		return false
	}
	return true
}

func getHash() hash.Hash {
	return sha256.New()
}
