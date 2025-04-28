package grandom

import (
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

// 生成UUID，去除其中的连字符
func GenerateUUIDWithoutHyphen() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

// RandomWait 在指定的最小和最大值范围内随机等待一段时间。
// 这个函数主要用于引入一个可配置的延迟，以增加操作的随机性和减少可预测性。
// 参数:
//
//	min - 随机等待时间的最小值（毫秒）。
//	max - 随机等待时间的最大值（毫秒）。
func RandomWait(min, max int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // 使用新的随机数生成器
	ms := r.Intn(max-min) + min                          // 生成一个随机数，范围在 [min, max] 之间
	time.Sleep(time.Duration(ms) * time.Millisecond)     // 等待随机毫秒数
}

// generateRandomNumber 生成指定长度的随机数字字符串。
// 参数 length 指定生成字符串的长度。
// 返回值为生成的随机数字字符串。
// 该函数适用于需要生成唯一标识符或随机数字序列的场景。
func GenerateRandomNumber(length int) string {
	// 定义可用的数字字符集。
	digits := "123456789"

	// 初始化结果切片，用于存储生成的随机数字。
	result := make([]byte, length)

	// 遍历结果切片的每个位置，填充随机选择的数字。
	for i := range result {
		// 从 digits 中随机选择一个字符。
		result[i] = digits[rand.Intn(len(digits))]
	}

	// 将结果切片转换为字符串并返回。
	return string(result)
}
