package server

import (
	"github.com/lithammer/dedent"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"testing"
	"time"
)

func TestLoadConfigFromFile(t *testing.T) {
	t.Run("should deserialize config correctly", func(t *testing.T) {
		file, err := os.CreateTemp("", "*")
		if err != nil {
			log.Fatal(err)
		}
		defer os.Remove(file.Name())
		file.Write([]byte(getConfigContent()))

		v, _ := os.ReadFile(file.Name())
		_ = v

		cfg, err := LoadConfigFromFile(file.Name())

		require.Nil(t, err)
		require.Equal(t, &Config{
			Address:          "0.0.0.0:3000",
			Difficulty:       26,
			SecretLength:     32,
			ChallengeTimeout: 10 * time.Second,
			Quotes:           []string{"quote1", "quote2"},
		}, cfg)

	})
}

func getConfigContent() string {
	return dedent.Dedent(`
		address: "0.0.0.0:3000"
		difficulty: 26
		secretLength: 32
		challengeTimeout: "10s"
		quotes: ["quote1", "quote2"]
	`)

}
