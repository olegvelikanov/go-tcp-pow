package pow

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"hash"
	"strconv"
	"time"
)

var ErrSolutionNotFound = fmt.Errorf("puzzle solution not found")

type Puzzle struct {
	Timestamp        time.Time
	CoveredPreImage  []byte
	CoveredBitsCount int
	Hash             []byte
}

type Solution struct {
	Timestamp time.Time
	PreImage  []byte
}

func NewPuzzle(bitsCount int, secret []byte) *Puzzle {
	timestamp := time.Now()
	preImage := calculatePreImage(timestamp, secret)
	hash := calculateHash(preImage)

	return &Puzzle{
		Timestamp:        timestamp,
		CoveredPreImage:  coverTrailingBits(preImage, bitsCount),
		CoveredBitsCount: bitsCount,
		Hash:             hash,
	}
}

func calculatePreImage(timestamp time.Time, secret []byte) []byte {
	return calculateHash([]byte(strconv.Itoa(timestamp.Nanosecond())), secret)
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
	preImage := make([]byte, len(p.CoveredPreImage))
	copy(preImage, p.CoveredPreImage)
	found := false

	n := 2 << p.CoveredBitsCount
	for i := 0; i < n; i++ {
		if bytes.Equal(p.Hash, calculateHash(preImage)) {
			found = true
			break
		}
		increment(preImage)
	}
	if !found {
		return nil, ErrSolutionNotFound
	}

	return &Solution{
		Timestamp: p.Timestamp,
		PreImage:  preImage,
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
	if !bytes.Equal(s.PreImage, calculatePreImage(s.Timestamp, secret)) {
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
