package pow

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
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

func NewPuzzle(bitsCount uint8, secret []byte) *Puzzle {
	timestamp := time.Now()
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

func coverTrailingBits(original []byte, bitsToReset int) []byte {
	result := make([]byte, len(original))
	idx := len(original) - 1 - (bitsToReset / 8)
	for i := range result {
		if i < idx {
			result[i] = original[i]
		} else if i > idx {
			result[i] = 0
		} else {
			result[i] = (original[i] >> (bitsToReset % 8)) << (bitsToReset % 8)
		}
	}
	return result
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

func (s *Solution) IsValid(secret []byte, timeout time.Duration) bool {
	if !bytes.Equal(s.HashOne, calculateHashOne(s.Timestamp, secret)) {
		return false
	}
	if s.Timestamp.Add(timeout).Before(time.Now()) {
		return false
	}
	return true
}

func getHash() hash.Hash {
	return sha256.New()
}
