package server

import (
	"github.com/benbjohnson/clock"
	"github.com/olegvelikanov/go-tcp-pow/internal/pkg/pow"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var (
	cl     = clock.NewMock()
	quotes = []string{"quote1", "quote2", "quote3"}
	app    = &App{
		challengeSecret:     []byte("secret"),
		challengeDifficulty: 5,
		challengeTimeout:    10 * time.Second,
		quotes:              quotes,
		clock:               cl,
	}
)

func init() {
	cl.Set(time.Now())
}

func TestApp_OnChallengeRequest(t *testing.T) {
	ts := time.Now()
	cl.Set(ts)

	puzzle := app.OnChallengeRequest()

	assert.Equal(t, uint8(5), puzzle.CoveredBitsCount)
	assert.Equal(t, ts, puzzle.Timestamp)
}

func TestApp_OnServiceRequest(t *testing.T) {
	t.Run("should grant service on correct solution", func(t *testing.T) {
		solution, err := app.OnChallengeRequest().Solve()
		assert.Nil(t, err)
		quote, err := app.OnServiceRequest(solution)
		require.Nil(t, err)
		require.Contains(t, quotes, string(quote))
	})

	t.Run("should return error on invalid solution", func(t *testing.T) {
		solution := &pow.Solution{
			Timestamp: time.Now(),
			HashOne:   []byte{},
		}
		quote, err := app.OnServiceRequest(solution)
		require.Nil(t, quote)
		require.Error(t, err, "invalid solution")
	})

	t.Run("should return error on timeout", func(t *testing.T) {
		solution, err := app.OnChallengeRequest().Solve()
		assert.Nil(t, err)

		cl.Set(time.Now().Add(time.Hour))
		quote, err := app.OnServiceRequest(solution)
		require.Nil(t, quote)
		require.Error(t, err, "solution timeout")
	})

}
