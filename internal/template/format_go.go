package template

import (
	"fmt"
	"strings"

	"github.com/nantokaworks/konst/internal/types"
	"github.com/nantokaworks/konst/internal/utils"
)

// ============================================================================
// 基本型フォーマット関数
// ============================================================================

// formatGo は、JSON の値を Go 用のリテラルに変換します。
// 対応: 文字列（日付の場合は time.Date(...) 形式）、数値、boolean、配列
func formatGo(value any) string {
	switch v := value.(type) {
	case string:
		return formatGoString(v)
	case float64:
		return formatGoFloat(v)
	case bool:
		return formatGoBool(v)
	case []interface{}:
		return formatGoSlice(v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// formatGoString は文字列値をフォーマットします（日付チェック込み）
func formatGoString(v string) string {
	if t, ok := tryParseDate(v); ok {
		year, month, day := t.Date()
		hour, min, sec := t.Clock()
		nsec := t.Nanosecond()
		return fmt.Sprintf("time.Date(%d, %s, %d, %d, %d, %d, %d, time.UTC)",
			year, utils.MonthConst(int(month)), day, hour, min, sec, nsec)
	}
	return fmt.Sprintf("%q", v)
}

// formatGoFloat は浮動小数点数をフォーマットします
func formatGoFloat(v float64) string {
	if v == float64(int(v)) {
		return fmt.Sprintf("%d", int(v))
	}
	return fmt.Sprintf("%f", v)
}

// formatGoBool はbool値をフォーマットします
func formatGoBool(v bool) string {
	if v {
		return "true"
	}
	return "false"
}

// ============================================================================
// 配列フォーマット関数
// ============================================================================

// formatGoSlice はスライスをフォーマットします
func formatGoSlice(v []any) string {
	if len(v) == 0 {
		return "nil"
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
		return formatGoNumberSlice(v)
	} else if allStrings {
		return formatGoStringSlice(v)
	} else if allBools {
		return formatGoBoolSlice(v)
	}

	// 混合型の場合
	var elems []string
	for _, elem := range v {
		elems = append(elems, formatGo(elem))
	}
	return "[]interface{}{" + strings.Join(elems, ", ") + "}"
}

// formatGoNumberSlice は数値スライスをフォーマットします
func formatGoNumberSlice(v []any) string {
	isInt := true
	var elems []string
	for _, elem := range v {
		num := elem.(float64)
		if num != float64(int(num)) {
			isInt = false
		}
		if isInt {
			elems = append(elems, fmt.Sprintf("%d", int(num)))
		} else {
			elems = append(elems, fmt.Sprintf("%f", num))
		}
	}
	if isInt {
		return "[]int{" + strings.Join(elems, ", ") + "}"
	}
	return "[]float64{" + strings.Join(elems, ", ") + "}"
}

// formatGoStringSlice は文字列スライスをフォーマットします
func formatGoStringSlice(v []any) string {
	var elems []string
	for _, elem := range v {
		elems = append(elems, fmt.Sprintf("%q", elem.(string)))
	}
	return "[]string{" + strings.Join(elems, ", ") + "}"
}

// formatGoBoolSlice はboolスライスをフォーマットします
func formatGoBoolSlice(v []any) string {
	var elems []string
	for _, elem := range v {
		elems = append(elems, fmt.Sprintf("%t", elem.(bool)))
	}
	return "[]bool{" + strings.Join(elems, ", ") + "}"
}

// ============================================================================
// 定数値フォーマット関数
// ============================================================================

// formatConstValue は、Definition の値を Go のコード形式にフォーマットします。
func formatConstValue(content any) string {
	def, ok := content.(types.Definition)
	if !ok {
		return formatGo(content)
	}

	switch def.Type {
	case types.DefinitionTypeInt, types.DefinitionTypeInt32, types.DefinitionTypeInt64,
		types.DefinitionTypeFloat, types.DefinitionTypeFloat32, types.DefinitionTypeFloat64:
		return formatGo(def.Value)
	case types.DefinitionTypeString:
		return formatGo(def.Value)
	case types.DefinitionTypeBool:
		return formatGo(def.Value)
	case types.DefinitionTypeDate:
		return formatGoDate(def)
	case types.DefinitionTypeTimestamp:
		return formatGoTimestamp(def)
	default:
		// 配列型の場合
		if strings.Contains(string(def.Type), "[]") {
			return formatGoArray(def)
		}
		return formatGo(def.Value)
	}
}

// formatGoDate は日付型の値をフォーマットします
func formatGoDate(def types.Definition) string {
	if def.GoMode == types.GoModeString {
		return formatGo(def.Value)
	}
	if def.GoMode == types.GoModeInt64 {
		if dateStr, ok := def.Value.(string); ok {
			if t, ok := tryParseDate(dateStr); ok {
				return fmt.Sprintf("%d", t.Unix())
			}
		}
		return formatGo(def.Value)
	}
	return formatGo(def.Value)
}

// formatGoTimestamp はtimestamp型の値をフォーマットします
func formatGoTimestamp(def types.Definition) string {
	if timestampStr, ok := def.Value.(string); ok {
		if t, ok := tryParseDate(timestampStr); ok {
			return fmt.Sprintf("%d", t.Unix())
		}
	}
	return formatGo(def.Value)
}

// formatGoArray は配列型の値をフォーマットします
func formatGoArray(def types.Definition) string {
	arrayValue, ok := def.Value.([]any)
	if !ok {
		return "nil"
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
			elements = append(elements, formatGoDate(tempDef))
		default:
			elements = append(elements, formatGo(elem))
		}
	}

	return "{" + strings.Join(elements, ", ") + "}"
}

// ============================================================================
// ヘルパー関数
// ============================================================================

// 汎用的なヘルパー関数は internal/utils/format_helpers.go に移動済み
