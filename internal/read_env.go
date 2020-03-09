package internal

import (
	"fmt"
	"github.com/joho/godotenv"
)

func ReadEnv(path string) (map[string]string, error) {
	envMap, err := godotenv.Read(path)
	if err != nil {
		return nil, fmt.Errorf("read: %w", err)
	}
	return envMap, nil
}
