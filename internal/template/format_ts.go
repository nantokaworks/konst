package template

import (
	"fmt"
	"strings"

	"github.com/nantokaworks/konst/internal/types"
)

// ============================================================================
// 基本型フォーマット関数
// ============================================================================

// formatTS は、JSON の値を TypeScript 用のリテラルに変換します。
// 対応: 文字列、数値、boolean、配列
func formatTS(value any) string {
	switch v := value.(type) {
	case string:
		return formatTSString(v)
	case float64:
		return formatTSNumber(v)
	case bool:
		return formatTSBool(v)
	case []interface{}:
		return formatTSArray(v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// formatTSString は文字列値をフォーマットします（日付チェック込み）
func formatTSString(v string) string {
	if t, ok := tryParseDate(v); ok {
		return fmt.Sprintf("new Date('%s')", t.Format("2006-01-02T15:04:05Z"))
	}
	return fmt.Sprintf("%q", v)
}

// formatTSNumber は数値をフォーマットします
func formatTSNumber(v float64) string {
	if v == float64(int(v)) {
		return fmt.Sprintf("%d", int(v))
	}
	return fmt.Sprintf("%f", v)
}

// formatTSBool はbool値をフォーマットします
func formatTSBool(v bool) string {
	if v {
		return "true"
	}
	return "false"
}

// ============================================================================
// 配列フォーマット関数
// ============================================================================

// formatTSArray は配列をTypeScript用にフォーマットします
func formatTSArray(v []any) string {
	if len(v) == 0 {
		return "[]"
	}

	// 型の統一性をチェック
	allNumbers, allStrings, allBools := true, true, true
	for _, elem := range v {
		if _, ok := elem.(float64); !ok {
			allNumbers = false
		}
		if _, ok := elem.(string); !ok {
			allStrings = false
		}
		if _, ok := elem.(bool); !ok {
			allBools = false
		}
	}

	if allNumbers {
		return formatTSNumberArray(v)
	} else if allStrings {
		return formatTSStringArray(v)
	} else if allBools {
		return formatTSBoolArray(v)
	}

	// 混合型の場合
	var elems []string
	for _, elem := range v {
		elems = append(elems, formatTS(elem))
	}
	return "[" + strings.Join(elems, ", ") + "]"
}

// formatTSNumberArray は数値配列をフォーマットします
func formatTSNumberArray(v []any) string {
	var elems []string
	for _, elem := range v {
		num := elem.(float64)
		if num == float64(int(num)) {
			elems = append(elems, fmt.Sprintf("%d", int(num)))
		} else {
			elems = append(elems, fmt.Sprintf("%f", num))
		}
	}
	return "[" + strings.Join(elems, ", ") + "]"
}

// formatTSStringArray は文字列配列をフォーマットします
func formatTSStringArray(v []any) string {
	var elems []string
	for _, elem := range v {
		str := elem.(string)
		if t, ok := tryParseDate(str); ok {
			elems = append(elems, fmt.Sprintf("new Date('%s')", t.Format("2006-01-02T15:04:05Z")))
		} else {
			elems = append(elems, fmt.Sprintf("%q", str))
		}
	}
	return "[" + strings.Join(elems, ", ") + "]"
}

// formatTSBoolArray はbool配列をフォーマットします
func formatTSBoolArray(v []any) string {
	var elems []string
	for _, elem := range v {
		elems = append(elems, fmt.Sprintf("%t", elem.(bool)))
	}
	return "[" + strings.Join(elems, ", ") + "]"
}

// ============================================================================
// 型変換関数
// ============================================================================

// 型変換関数は internal/utils/format_helpers.go に移動済み

// ============================================================================
// 定数値フォーマット関数（Definition対応）
// ============================================================================

// formatTSConstValue は、Definition の値を TypeScript のコード形式にフォーマットします。
func formatTSConstValue(content any) string {
	def, ok := content.(types.Definition)
	if !ok {
		return formatTS(content)
	}

	switch def.Type {
	case types.DefinitionTypeInt, types.DefinitionTypeInt32, types.DefinitionTypeInt64,
		types.DefinitionTypeFloat, types.DefinitionTypeFloat32, types.DefinitionTypeFloat64:
		if num, ok := def.Value.(float64); ok {
			return formatTSNumber(num)
		}
		return formatTS(def.Value)
	case types.DefinitionTypeString:
		if str, ok := def.Value.(string); ok {
			return formatTSString(str)
		}
		return formatTS(def.Value)
	case types.DefinitionTypeBool:
		if b, ok := def.Value.(bool); ok {
			return formatTSBool(b)
		}
		return formatTS(def.Value)
	case types.DefinitionTypeDate:
		return formatTSDate(def)
	case types.DefinitionTypeTimestamp:
		return formatTSTimestamp(def)
	default:
		// 配列型の場合
		if strings.Contains(string(def.Type), "[]") {
			return formatTSDefinitionArray(def)
		}
		return formatTS(def.Value)
	}
}

// formatTSDate は日付型の値をTypeScript用にフォーマットします
func formatTSDate(def types.Definition) string {
	switch def.TSMode {
	case types.ModeString:
		if str, ok := def.Value.(string); ok {
			return formatTSString(str)
		}
		return formatTS(def.Value)
	case types.ModeDate:
		if dateStr, ok := def.Value.(string); ok {
			if t, ok := tryParseDate(dateStr); ok {
				return fmt.Sprintf("new Date('%s')", t.Format("2006-01-02T15:04:05Z"))
			}
		}
		return formatTS(def.Value)
	case types.ModeBigInt:
		if dateStr, ok := def.Value.(string); ok {
			if t, ok := tryParseDate(dateStr); ok {
				return fmt.Sprintf("BigInt(%d)", t.Unix())
			}
		}
		return formatTS(def.Value)
	default:
		return formatTS(def.Value)
	}
}

// formatTSTimestamp はtimestamp型の値をTypeScript用にフォーマットします
func formatTSTimestamp(def types.Definition) string {
	if timestampStr, ok := def.Value.(string); ok {
		if t, ok := tryParseDate(timestampStr); ok {
			return fmt.Sprintf("%d", t.Unix())
		}
	}
	return formatTS(def.Value)
}

// formatTSDefinitionArray は配列型の値をTypeScript用にフォーマットします
func formatTSDefinitionArray(def types.Definition) string {
	arrayValue, ok := def.Value.([]any)
	if !ok {
		return "[]"
	}

	var elements []string
	baseType := strings.TrimSuffix(string(def.Type), "[]")

	for _, elem := range arrayValue {
		switch baseType {
		case "date":
			tempDef := types.Definition{
				Type:   types.DefinitionTypeDate,
				Value:  elem,
				TSMode: def.TSMode,
				GoMode: def.GoMode,
			}
			elements = append(elements, formatTSDate(tempDef))
		default:
			elements = append(elements, formatTS(elem))
		}
	}

	return "[" + strings.Join(elements, ", ") + "]"
}
