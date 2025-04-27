package gstrutils

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// 适用于处理复杂的嵌套数据结构，能够灵活地根据键路径查找值
func findMapByKey(sourceMap map[string]any, keys []string) any {
	if len(sourceMap) <= 0 {
		return nil
	}

	if len(keys) <= 0 {
		return sourceMap
	}

	k := keys[0]
	next, ok := sourceMap[k]
	if !ok {
		return nil
	}
	if len(keys) == 1 {
		return next
	}
	switch nextMap := next.(type) {
	case map[string]any:
		return findMapByKey(nextMap, keys[1:])
	case []any:
		newMap := make(map[string]any, len(nextMap))
		for i, ele := range nextMap {
			newMap[strconv.Itoa(i)] = ele
		}
		return findMapByKey(newMap, keys[1:])
	default:
		log.Println("findMapByKey: unknown type")
		return nil
	}
}

// 变量替换
func ValReplace(text string, sourceMap map[string]any) (string, error) {

	return os.Expand(text, func(completeKey string) string {
		keys := strings.Split(completeKey, ".")
		val := findMapByKey(sourceMap, keys)
		var valStr string
		switch act := val.(type) {
		case int:
			valStr = strconv.FormatInt(int64(act), 10)
		case int8:
			valStr = strconv.FormatInt(int64(act), 10)
		case int16:
			valStr = strconv.FormatInt(int64(act), 10)
		case int32:
			valStr = strconv.FormatInt(int64(act), 10)
		case int64:
			valStr = strconv.FormatInt(act, 10)
		case uint:
			valStr = strconv.FormatInt(int64(act), 10)
		case uint8:
			valStr = strconv.FormatInt(int64(act), 10)
		case uint16:
			valStr = strconv.FormatInt(int64(act), 10)
		case uint32:
			valStr = strconv.FormatInt(int64(act), 10)
		case uint64:
			valStr = strconv.FormatInt(int64(act), 10)
		case float64:
			valStr = strconv.FormatFloat(act, 'f', -1, 64)
		case float32:
			valStr = strconv.FormatFloat(float64(act), 'f', -1, 32)
		case bool:
			valStr = strconv.FormatBool(act)
		case string:
			valStr = act
		}
		return valStr
	}), nil
}
