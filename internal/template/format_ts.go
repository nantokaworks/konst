package template

import (
	"fmt"
	"strings"
	"time"

	"github.com/nantokaworks/konst/internal/types"
)

// formatTS は、JSON の値を TypeScript 用のリテラルに変換します。
// DefinitionContent（または map[string]interface{}）にも対応します。
func formatTS(value interface{}) string {
	// ① map[string]interface{} の場合
	if m, ok := value.(map[string]interface{}); ok {
		if v, exists := m["value"]; exists {
			if t, hasType := m["type"].(string); hasType {
				// tsMode を string 経由で取得しているケースの場合は、enum に変換する
				var tsMode types.TSMode
				if mode, ok := m["tsMode"].(string); ok && mode != "" {
					tsMode = types.TSMode(mode)
				}
				switch t {
				case "int64", "uint64":
					if tsMode == types.ModeNumber {
						if num, ok := v.(float64); ok {
							return fmt.Sprintf("%d as const", int64(num))
						}
					} else {
						// TSModeBigInt またはその他の場合
						if num, ok := v.(float64); ok {
							return fmt.Sprintf("%dn as const", int64(num))
						}
					}
				case "int", "float":
					if num, ok := v.(float64); ok {
						if num == float64(int(num)) {
							return fmt.Sprintf("%d as const", int(num))
						}
						return fmt.Sprintf("%f as const", num)
					}
				case "date":
					{
						var dateMode types.DateMode
						if dm, ok := m["mode"].(string); ok && dm != "" {
							dateMode = types.DateMode(dm)
						}
						if dateMode == types.DateModeString {
							return fmt.Sprintf("%q as const", v)
						}
						return fmt.Sprintf("new Date(%q)", v)
					}
				default:
					return fmt.Sprintf("%q as const", v)
				}
			}
			return formatTS(v)
		}
	}
	// ② Definition 型の場合
	if d, ok := value.(types.Definition); ok {
		v := d.Value
		t := d.Type
		tsMode := d.TSMode
		switch t {
		case "int64", "uint64":
			if num, ok := v.(float64); ok {
				if tsMode == types.ModeNumber {
					return fmt.Sprintf("%d as const", int64(num))
				}
				return fmt.Sprintf("%dn as const", int64(num))
			}
		case "int", "float":
			if num, ok := v.(float64); ok {
				if num == float64(int(num)) {
					return fmt.Sprintf("%d as const", int(num))
				}
				return fmt.Sprintf("%f as const", num)
			}
		case "date":
			if d.DateMode == string(types.DateModeString) {
				return fmt.Sprintf("%q as const", v)
			}
			return fmt.Sprintf("new Date(%q)", v)
		default:
			return fmt.Sprintf("%q as const", v)
		}
	}
	// ポインター型の Definition への対応
	if d, ok := value.(*types.Definition); ok {
		return formatTS(*d)
	}
	// ④ その他通常の型の処理
	switch v := value.(type) {
	case string:
		if t, ok := tryParseDate(v); ok {
			return fmt.Sprintf("new Date(%q)", t.Format(time.RFC3339))
		}
		return fmt.Sprintf("%q as const", v)
	case float64:
		if v == float64(int(v)) {
			return fmt.Sprintf("%d as const", int(v))
		}
		return fmt.Sprintf("%f as const", v)
	case bool:
		if v {
			return "true as const"
		}
		return "false as const"
	case []interface{}:
		var elems []string
		for _, elem := range v {
			elems = append(elems, formatTS(elem))
		}
		return "[" + strings.Join(elems, ", ") + "] as const"
	default:
		return fmt.Sprintf("%v as const", v)
	}
}

// convertTSType は、Go の型名を TypeScript の型名に変換します。
func convertTSType(goType string) string {
	switch goType {
	case "int", "int64", "float", "float64", "uint64":
		return "number"
	case "string":
		return "string"
	case "bool":
		return "boolean"
	default:
		return goType
	}
}
