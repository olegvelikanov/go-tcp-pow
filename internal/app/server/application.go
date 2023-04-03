package server

import (
	cryptorand "crypto/rand"
	"fmt"
	"github.com/olegvelikanov/go-tcp-pow/internal/pkg/pow"
	"math/rand"
	"time"
)

type Application interface {
	onChallengeRequest() *pow.Puzzle

	onServiceRequest(solution *pow.Solution) ([]byte, error)
}

type App struct {
	challengeSecret     []byte
	challengeDifficulty uint8
	challengeTimeout    time.Duration
	quotes              []string
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
	}, nil
}

func (w *App) onChallengeRequest() *pow.Puzzle {
	return pow.NewPuzzle(w.challengeDifficulty, w.challengeSecret)
}

func (w *App) onServiceRequest(solution *pow.Solution) ([]byte, error) {
	if !solution.IsValid(w.challengeSecret, w.challengeTimeout) {
		return nil, fmt.Errorf("invalid solution")
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
