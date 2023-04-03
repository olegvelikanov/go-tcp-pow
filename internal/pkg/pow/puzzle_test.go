package pow

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var (
	secret   = generateSecret()
	duration = 60 * time.Second
)

func TestPuzzle_Solve(t *testing.T) {
	t.Run("should correctly solve puzzle", func(t *testing.T) {
		puzzle := NewPuzzle(24, secret)
		solution, err := puzzle.Solve()
		assert.Nil(t, err)
		require.True(t, solution.IsValid(secret, duration))
	})

	t.Run("should return error for unsolvable puzzle", func(t *testing.T) {
		puzzle := NewPuzzle(24, secret)
		puzzle.HashTwo = make([]byte, len(puzzle.HashTwo))
		solution, err := puzzle.Solve()
		require.Nil(t, solution)
		assert.Errorf(t, err, "puzzle solution not found")
	})
}

func TestSolution_IsValid(t *testing.T) {
	timestamp := time.Now()

	t.Run("should return true for valid solution", func(t *testing.T) {
		solution := &Solution{
			Timestamp: timestamp,
			HashOne:   calculateHashOne(timestamp, []byte("secret")),
		}
		assert.True(t, solution.IsValid([]byte("secret"), duration))
	})

	t.Run("should return true if timestamps in different timezones", func(t *testing.T) {
		loc1 := time.FixedZone("loc1", 0)
		loc2 := time.FixedZone("loc2", 3)
		solution := &Solution{
			Timestamp: timestamp.In(loc1),
			HashOne:   calculateHashOne(timestamp.In(loc2), []byte("secret")),
		}
		assert.True(t, solution.IsValid([]byte("secret"), duration))
	})

	t.Run("should return false for invalid solution", func(t *testing.T) {
		solution := &Solution{
			Timestamp: timestamp,
			HashOne:   increment(calculateHashOne(timestamp, []byte("secret"))),
		}
		assert.False(t, solution.IsValid([]byte("secret"), duration))
	})

	t.Run("should return false on timeout", func(t *testing.T) {
		solution := &Solution{
			Timestamp: timestamp.Add(24 * time.Hour),
			HashOne:   increment(calculateHashOne(timestamp, []byte("secret"))),
		}
		assert.False(t, solution.IsValid([]byte("secret"), duration))
	})

}

func Test_resetTrailingBits_(t *testing.T) {
	tests := []struct {
		original  string
		bitsCount int
		expected  string
	}{
		{"1f", 4, "10"},
		{"1f", 5, "00"},
		{"abffffff", 23, "ab800000"},
		{"abffffff", 24, "ab000000"},
		{"abffffff", 25, "aa000000"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("should correctly reset %d trailing bits from #%s", test.bitsCount, test.original), func(t *testing.T) {
			assert.Equal(t, decodeHex(test.expected), coverTrailingBits(decodeHex(test.original), test.bitsCount))
		})
	}
}

func Test_increment(t *testing.T) {
	tests := []struct {
		original string
		expected string
	}{
		{"00", "01"},
		{"0c", "0d"},
		{"0f", "10"},
		{"10", "11"},
		{"ff", "00"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("should correctly find increment for #%s", test.original), func(t *testing.T) {
			assert.Equal(t, decodeHex(test.expected), increment(decodeHex(test.original)))
		})
	}
}

func decodeHex(s string) []byte {
	result, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return result
}

func generateSecret() []byte {
	result := make([]byte, 64)
	rand.Read(result)
	return result
}
