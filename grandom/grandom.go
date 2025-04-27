package grandom

import (
	"strings"

	"github.com/google/uuid"
)

// 生成UUID，去除其中的连字符
func GenerateUUIDWithoutHyphen() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
