package server

import (
	"fmt"
	"github.com/olegvelikanov/go-tcp-pow/internal/pkg/pow"
	"math/rand"
	"time"
)

type Application interface {
	onChallengeRequest() *pow.Puzzle

	onServiceRequest(solution *pow.Solution) ([]byte, error)
}

type WordOfWisdomApp struct {
	challengeSecret     []byte
	challengeDifficulty uint8
	challengeTimeout    time.Duration
	quotes              []string
}

func NewWordOfWisdomApp() *WordOfWisdomApp {
	return &WordOfWisdomApp{
		challengeSecret:     []byte("secret"),
		challengeDifficulty: 24,
		challengeTimeout:    10 * time.Second,
		quotes:              []string{"quote1", "quote2"},
	}
}

func (w *WordOfWisdomApp) onChallengeRequest() *pow.Puzzle {
	return pow.NewPuzzle(w.challengeDifficulty, w.challengeSecret)
}

func (w *WordOfWisdomApp) onServiceRequest(solution *pow.Solution) ([]byte, error) {
	if !solution.IsValid(w.challengeSecret, w.challengeTimeout) {
		return nil, fmt.Errorf("invalid solution")
	}

	return []byte(w.pickRandomQuote()), nil
}

func (w *WordOfWisdomApp) pickRandomQuote() string {
	return w.quotes[rand.Intn(len(w.quotes))]
}
