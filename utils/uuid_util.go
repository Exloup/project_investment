package utils

import "github.com/google/uuid"

func UUIdGenerate() string {
	return uuid.New().String()
}
