// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import (
	"math/rand"
	"time"
)

func getRandomId() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	bytes := make([]byte, 10)
	for i := range bytes {
		bytes[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(bytes)
}

type Config struct {
	Path        string `config:"path"`
	Type        string `config:"type"`
	RunId       string `config:"runid"`
	Environment string `config:"environment"`
	Project     string `config:"project"`
	Link        string `config:"link"`
}

var DefaultConfig = Config{
	Path:        ".",
	Type:        "junit",
	RunId:       getRandomId(),
	Environment: "",
	Project:     "",
	Link:        "",
}
