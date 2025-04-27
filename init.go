package gtools

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateUUIDWithoutHyphen() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
