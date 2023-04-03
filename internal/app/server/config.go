package server

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type Config struct {
	Address          string        `yaml:"address"`
	Difficulty       uint8         `yaml:"difficulty"`
	SecretLength     int           `yaml:"secretLength"`
	ChallengeTimeout time.Duration `yaml:"challengeTimeout"`
	Quotes           []string      `yaml:"quotes"`
}

func LoadConfigFromFile(name string) (*Config, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("opening file: %s", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	config := &Config{}
	err = decoder.Decode(config)
	if err != nil {
		return nil, fmt.Errorf("decoding config: %s", err)
	}
	return config, nil

}
