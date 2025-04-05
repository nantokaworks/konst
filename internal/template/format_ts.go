package template

import (
	"fmt"
	"strings"
	"time"

	"github.com/nantokaworks/konst/internal/types"
)

func formatTS(value interface{}) string {
	// 数値フォーマット用のヘルパー関数
	formatNumeric := func(num float64, tsMode types.TSMode) string {
		switch tsMode {
		case types.ModeNumber:
			return fmt.Sprintf("%d as const", int64(num))
		case types.ModeBigInt:
			return fmt.Sprintf("%dn as const", int64(num))
		default:
			// デフォルトは数値として出力
			return fmt.Sprintf("%d as const", int64(num))
		}
	}

	// ① map[string]interface{} の場合
	if m, ok := value.(map[string]interface{}); ok {
		if v, exists := m["value"]; exists {
			if tStr, hasType := m["type"].(string); hasType {
				t := types.DefinitionType(tStr)
				var tsMode types.TSMode
				if mode, ok := m["tsMode"].(string); ok && mode != "" {
					tsMode = types.TSMode(mode)
				}
				switch t {
				case types.DefinitionTypeInt64, types.DefinitionTypeUint64, types.DefinitionTypeFloat64:
					if num, ok := v.(float64); ok {
						return formatNumeric(num, tsMode)
					}
				case types.DefinitionTypeUint32, types.DefinitionTypeUint:
					if num, ok := v.(float64); ok {
						return formatNumeric(num, tsMode)
					}
				case types.DefinitionTypeInt, types.DefinitionTypeFloat, types.DefinitionTypeInt32, types.DefinitionTypeFloat32:
					if num, ok := v.(float64); ok {
						if num == float64(int(num)) {
							return fmt.Sprintf("%d as const", int(num))
						}
						return fmt.Sprintf("%f as const", num)
					}
				case types.DefinitionTypeDate:
					switch tsMode {
					case types.ModeString:
						return fmt.Sprintf("%q as const", v)
					case types.ModeDate:
						return fmt.Sprintf("new Date(%q)", v)
					default:
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
		if strings.HasSuffix(string(t), "[]") {
			if arr, ok := v.([]interface{}); ok {
				var elems []string
				for _, elem := range arr {
					elems = append(elems, formatTS(elem))
				}
				return "[" + strings.Join(elems, ", ") + "] as const"
			}
			return formatTS(v)
		}
		switch t {
		case types.DefinitionTypeInt64, types.DefinitionTypeUint64, types.DefinitionTypeFloat64:
			if num, ok := v.(float64); ok {
				return formatNumeric(num, tsMode)
			}
		case types.DefinitionTypeUint32, types.DefinitionTypeUint:
			if num, ok := v.(float64); ok {
				return formatNumeric(num, tsMode)
			}
		case types.DefinitionTypeInt, types.DefinitionTypeFloat, types.DefinitionTypeInt32, types.DefinitionTypeFloat32:
			if num, ok := v.(float64); ok {
				if num == float64(int(num)) {
					return fmt.Sprintf("%d as const", int(num))
				}
				return fmt.Sprintf("%f as const", num)
			}
		case types.DefinitionTypeDate:
			switch tsMode {
			case types.ModeString:
				return fmt.Sprintf("%q as const", v)
			case types.ModeDate:
				return fmt.Sprintf("new Date(%q)", v)
			default:
				return fmt.Sprintf("new Date(%q)", v)
			}
		case types.DefinitionTypeBool:
			if b, ok := v.(bool); ok {
				if b {
					return "true as const"
				}
				return "false as const"
			}
			return fmt.Sprintf("%v as const", v)
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
