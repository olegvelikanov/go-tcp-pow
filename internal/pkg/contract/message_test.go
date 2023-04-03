package contract

import (
	"crypto/rand"
	"github.com/olegvelikanov/go-tcp-pow/internal/pkg/pow"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestDeserialize_ChallengeRequest(t *testing.T) {

	t.Run("should serialize and deserialize ChallengeRequest symmetrically", func(t *testing.T) {
		original := &ChallengeRequest{}
		bytes := make([]byte, 1)

		n, err := Serialize(original, bytes)
		assert.Nil(t, err)
		require.Equal(t, 1, n)
		deserialized, err := Deserialize(bytes)
		require.Nil(t, err)
		require.Equal(t, original, deserialized)
	})
}

func TestDeserialize_ChallengeResponse(t *testing.T) {

	t.Run("should serialize and deserialize ChallengeResponse symmetrically", func(t *testing.T) {
		original := &ChallengeResponse{
			Puzzle: &pow.Puzzle{
				Timestamp:        time.Now().In(time.UTC),
				CoveredHashOne:   generateBytes(32),
				CoveredBitsCount: 24,
				HashTwo:          generateBytes(32),
			},
		}
		bytes := make([]byte, 82)

		n, err := Serialize(original, bytes)
		assert.Nil(t, err)
		require.Equal(t, 82, n)
		deserialized, err := Deserialize(bytes)
		require.Nil(t, err)
		require.Equal(t, original, deserialized)
	})
}

func TestDeserialize_ServiceRequest(t *testing.T) {

	t.Run("should serialize and deserialize ServiceRequest symmetrically", func(t *testing.T) {
		original := &ServiceRequest{
			PuzzleSolution: &pow.Solution{
				Timestamp: time.Now().In(time.UTC),
				HashOne:   generateBytes(32),
			},
		}
		bytes := make([]byte, 49)

		n, err := Serialize(original, bytes)
		assert.Nil(t, err)
		require.Equal(t, 49, n)
		deserialized, err := Deserialize(bytes)
		require.Nil(t, err)
		require.Equal(t, original, deserialized)
	})
}

func TestDeserialize_ServiceResponse(t *testing.T) {

	t.Run("should serialize and deserialize ServiceResponse symmetrically", func(t *testing.T) {
		original := &ServiceResponse{
			Quote: generateBytes(32),
		}
		bytes := make([]byte, 41)

		n, err := Serialize(original, bytes)
		assert.Nil(t, err)
		require.Equal(t, 41, n)
		deserialized, err := Deserialize(bytes)
		require.Nil(t, err)
		require.Equal(t, original, deserialized)
	})
}

func generateBytes(n int) []byte {
	result := make([]byte, n)
	rand.Read(result)
	return result
}
