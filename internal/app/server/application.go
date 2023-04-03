package server

import (
	cryptorand "crypto/rand"
	"fmt"
	"github.com/benbjohnson/clock"
	"github.com/olegvelikanov/go-tcp-pow/internal/pkg/pow"
	"math/rand"
	"time"
)

type Application interface {
	OnChallengeRequest() *pow.Puzzle

	OnServiceRequest(solution *pow.Solution) ([]byte, error)
}

type App struct {
	challengeSecret     []byte
	challengeDifficulty uint8
	challengeTimeout    time.Duration
	quotes              []string
	clock               clock.Clock
}

func NewApp(config *Config) (*App, error) {
	secret, err := generateSecret(config.SecretLength)
	if err != nil {
		return nil, fmt.Errorf("generating secret: %s", err)
	}
	return &App{
		challengeSecret:     secret,
		challengeDifficulty: config.Difficulty,
		challengeTimeout:    config.ChallengeTimeout,
		quotes:              config.Quotes,
		clock:               clock.New(),
	}, nil
}

func (w *App) OnChallengeRequest() *pow.Puzzle {
	return pow.NewPuzzle(w.challengeDifficulty, w.challengeSecret, w.clock)
}

func (w *App) OnServiceRequest(solution *pow.Solution) ([]byte, error) {
	if !solution.IsValid(w.challengeSecret) {
		return nil, fmt.Errorf("invalid solution")
	}

	if solution.Timestamp.Add(w.challengeTimeout).Before(w.clock.Now()) {
		return nil, fmt.Errorf("solution timeout")
	}

	return []byte(w.pickRandomQuote()), nil
}

func (w *App) pickRandomQuote() string {
	return w.quotes[rand.Intn(len(w.quotes))]
}

func generateSecret(n int) ([]byte, error) {
	result := make([]byte, n)
	_, err := cryptorand.Read(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
