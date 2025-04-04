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
					if tsMode == types.TSModeNumber {
						if num, ok := v.(float64); ok {
							return fmt.Sprintf("%d", int64(num))
						}
					} else {
						// TSModeBigInt またはその他の場合
						if num, ok := v.(float64); ok {
							return fmt.Sprintf("%dn", int64(num))
						}
					}
				case "int", "float":
					if num, ok := v.(float64); ok {
						if num == float64(int(num)) {
							return fmt.Sprintf("%d", int(num))
						}
						return fmt.Sprintf("%f", num)
					}
				case "date":
					// TSMode により出力を分岐
					if tsMode == types.TSModeString {
						// 文字列として出力
						return fmt.Sprintf("%q", v)
					} else {
						// Date 型として new Date(...) を出力
						return fmt.Sprintf("new Date(%q)", v)
					}
				default:
					return fmt.Sprintf("%q", v)
				}
			}
			return formatTS(v)
		}
	}
	// ② DefinitionContent 型（構造体の場合）
	if d, ok := value.(types.DefinitionContent); ok {
		if d.ConstContent != nil {
			v := d.ConstContent.Value
			t := d.ConstContent.Type
			tsMode := d.ConstContent.TSMode // TSMode 型
			switch t {
			case "int64", "uint64":
				if tsMode != "" {
					if tsMode == types.TSModeNumber {
						if num, ok := v.(float64); ok {
							return fmt.Sprintf("%d", int64(num))
						}
					} else { // TSModeBigInt またはそれ以外
						if num, ok := v.(float64); ok {
							return fmt.Sprintf("%dn", int64(num))
						}
					}
				} else {
					if num, ok := v.(float64); ok {
						return fmt.Sprintf("%dn", int64(num))
					}
				}
			case "int", "float":
				if num, ok := v.(float64); ok {
					if num == float64(int(num)) {
						return fmt.Sprintf("%d", int(num))
					}
					return fmt.Sprintf("%f", num)
				}
			case "date":
				if tsMode == types.TSModeString {
					return fmt.Sprintf("%q", v)
				}
				return fmt.Sprintf("new Date(%q)", v)
			default:
				return fmt.Sprintf("%q", v)
			}
		}
	}
	// ③ ポインター型の DefinitionContent への対応
	if d, ok := value.(*types.DefinitionContent); ok {
		if d.ConstContent != nil {
			v := d.ConstContent.Value
			t := d.ConstContent.Type
			tsMode := d.ConstContent.TSMode
			switch t {
			case "int64", "uint64":
				if tsMode != "" {
					if tsMode == types.TSModeNumber {
						if num, ok := v.(float64); ok {
							return fmt.Sprintf("%d", int64(num))
						}
					} else {
						if num, ok := v.(float64); ok {
							return fmt.Sprintf("%dn", int64(num))
						}
					}
				} else {
					if num, ok := v.(float64); ok {
						return fmt.Sprintf("%dn", int64(num))
					}
				}
			case "int", "float":
				if num, ok := v.(float64); ok {
					if num == float64(int(num)) {
						return fmt.Sprintf("%d", int(num))
					}
					return fmt.Sprintf("%f", num)
				}
			case "date":
				if tsMode == types.TSModeString {
					return fmt.Sprintf("%q", v)
				}
				return fmt.Sprintf("new Date(%q)", v)
			default:
				return fmt.Sprintf("%q", v)
			}
		}
	}
	// ④ その他通常の型の処理
	switch v := value.(type) {
	case string:
		if t, ok := tryParseDate(v); ok {
			return fmt.Sprintf("new Date(%q)", t.Format(time.RFC3339))
		}
		return fmt.Sprintf("%q", v)
	case float64:
		if v == float64(int(v)) {
			return fmt.Sprintf("%d", int(v))
		}
		return fmt.Sprintf("%f", v)
	case bool:
		if v {
			return "true"
		}
		return "false"
	case []interface{}:
		var elems []string
		for _, elem := range v {
			elems = append(elems, formatTS(elem))
		}
		return "[" + strings.Join(elems, ", ") + "]"
	default:
		return fmt.Sprintf("%v", v)
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
