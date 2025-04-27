package gtools

import (
	"strings"

	"github.com/google/uuid"
)

var AAA = new(Aaa)

type Aaa struct {
}

func (aaa *Aaa) GenerateUUIDWithoutHyphen() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
