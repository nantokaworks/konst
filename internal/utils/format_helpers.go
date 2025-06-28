package utils

import (
	"fmt"
	"strings"

	"github.com/nantokaworks/konst/internal/types"
)

// AsString は任意の値を文字列に変換します
func AsString(v any) string {
	return fmt.Sprintf("%v", v)
}

// MonthConst は月番号を Go の time.Month 定数に変換します。
func MonthConst(m int) string {
	if m < 1 || m > 12 {
		return "time.January"
	}
	months := []string{
		"", "time.January", "time.February", "time.March", "time.April", "time.May", "time.June",
		"time.July", "time.August", "time.September", "time.October", "time.November", "time.December",
	}
	return months[m]
}

// HasDate は定義の中に日付型があるかチェックします
func HasDate(defs map[string]types.Definition) bool {
	if defs == nil {
		return false
	}
	for _, def := range defs {
		if def.Type == "date" {
			return true
		}
	}
	return false
}

// ConvertTSType は Go の型名を TypeScript の型名に変換します。
func ConvertTSType(goType string) string {
	switch goType {
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64":
		return "number"
	case "string":
		return "string"
	case "bool":
		return "boolean"
	case "date":
		return "Date"
	case "timestamp":
		return "number" // Unix timestamp
	default:
		// 配列型の処理
		if strings.HasSuffix(goType, "[]") {
			baseType := strings.TrimSuffix(goType, "[]")
			return ConvertTSType(baseType) + "[]"
		}
		return "any"
	}
}
