package template

import (
	"fmt"
	"strings"

	"github.com/nantokaworks/konst/internal/types"
)

// formatGo は、JSON の値を Go 用のリテラルに変換します。
// 対応: 文字列（日付の場合は time.Date(...) 形式）、数値、boolean、配列
func formatGo(value interface{}) string {
	// 数値フォーマットのヘルパー関数
	formatNumber := func(v float64, isInt bool) string {
		if isInt {
			return fmt.Sprintf("%d", int(v))
		}
		return fmt.Sprintf("%f", v)
	}

	switch v := value.(type) {
	case string:
		if t, ok := tryParseDate(v); ok {
			year, month, day := t.Date()
			hour, min, sec := t.Clock()
			nsec := t.Nanosecond()
			return fmt.Sprintf("time.Date(%d, %s, %d, %d, %d, %d, %d, time.UTC)",
				year, monthConst(int(month)), day, hour, min, sec, nsec)
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
		if len(v) == 0 {
			return "nil"
		}
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
			isInt := true
			var elems []string
			for _, elem := range v {
				num := elem.(float64)
				if num != float64(int(num)) {
					isInt = false
				}
				elems = append(elems, formatNumber(num, isInt))
			}
			if isInt {
				return "[]int{" + strings.Join(elems, ", ") + "}"
			}
			return "[]float64{" + strings.Join(elems, ", ") + "}"
		} else if allStrings {
			var elems []string
			for _, elem := range v {
				elems = append(elems, fmt.Sprintf("%q", elem.(string)))
			}
			return "[]string{" + strings.Join(elems, ", ") + "}"
		} else if allBools {
			var elems []string
			for _, elem := range v {
				elems = append(elems, fmt.Sprintf("%t", elem.(bool)))
			}
			return "[]bool{" + strings.Join(elems, ", ") + "}"
		}
		return fmt.Sprintf("%#v", value)
	case types.Definition:
		val := v.Value
		typ := v.Type // DefinitionType
		// 追加: 配列型の場合は value 部分を再帰的に処理する
		if strings.HasSuffix(string(typ), "[]") {
			return formatGo(val)
		}
		switch typ {
		case types.DefinitionTypeInt64, types.DefinitionTypeUint64, types.DefinitionTypeUint, types.DefinitionTypeFloat64:
			if num, ok := val.(float64); ok {
				return fmt.Sprintf("%d", int64(num))
			}
		case types.DefinitionTypeUint32:
			if num, ok := val.(float64); ok {
				return fmt.Sprintf("%d", uint32(num))
			}
		case types.DefinitionTypeInt, types.DefinitionTypeFloat, types.DefinitionTypeInt32, types.DefinitionTypeFloat32:
			if num, ok := val.(float64); ok {
				if num == float64(int(num)) {
					return fmt.Sprintf("%d", int(num))
				}
				return fmt.Sprintf("%f", num)
			}
		case types.DefinitionTypeDate:
			switch v.GoMode {
			case types.GoModeString:
				return fmt.Sprintf("%q", val)
			case types.GoModeInt, types.GoModeInt64:
				if t, ok := tryParseDate(val.(string)); ok {
					return fmt.Sprintf("%d", t.Unix())
				}
				return fmt.Sprintf("0 /* invalid date: %q */", val)
			case types.GoModeTimestamp: // TSMode == "timestamp"
				if t, ok := tryParseDate(val.(string)); ok {
					return fmt.Sprintf("%d", t.UnixNano()/1e6)
				}
				return fmt.Sprintf("0 /* invalid date: %q */", val)
			default:
				if t, ok := tryParseDate(val.(string)); ok {
					year, month, day := t.Date()
					hour, min, sec := t.Clock()
					nsec := t.Nanosecond()
					return fmt.Sprintf("time.Date(%d, %s, %d, %d, %d, %d, %d, time.UTC)",
						year, monthConst(int(month)), day, hour, min, sec, nsec)
				}
				// フォールバック: 無効な日付文字列の場合
				return fmt.Sprintf("time.Now() /* invalid date: %q */", val)
			}
		case types.DefinitionTypeBool:
			if b, ok := val.(bool); ok {
				return fmt.Sprintf("%t", b)
			}
			return fmt.Sprintf("%v", val)
		default:
			return fmt.Sprintf("%q", val)
		}
	case *types.Definition:
		return formatGo(*v)
	default:
		return fmt.Sprintf("%v", value)
	}

	return ""
}

// monthConst は、月番号から Go の time.Month 定数名を返します。
func monthConst(m int) string {
	months := []string{
		"time.January", "time.February", "time.March", "time.April",
		"time.May", "time.June", "time.July", "time.August",
		"time.September", "time.October", "time.November", "time.December",
	}
	if m >= 1 && m <= 12 {
		return months[m-1]
	}
	return fmt.Sprintf("time.Month(%d)", m)
}

// formatConstValue は、定数の場合、content 内の "value" キーからリテラル値を抽出してフォーマットします。
// こちらは Go 用の定数出力用です。
func formatConstValue(content interface{}) string {
	// ① map[string]interface{} として扱える場合
	if m, ok := content.(map[string]interface{}); ok {
		if v, exists := m["value"]; exists {
			return formatGo(v)
		}
		return formatGo(m)
	}
	// ② Definition 型の場合 (旧 DefinitionContent)
	if d, ok := content.(types.Definition); ok {
		return formatGo(d)
	} else if d, ok := content.(*types.Definition); ok {
		return formatGo(d)
	}
	// ③ その他の場合は、直接 formatGo を呼ぶ
	return formatGo(content)
}

func asString(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

func hasDate(defs map[string]types.Definition) bool {
	if defs == nil {
		return false
	}
	for _, def := range defs {
		// def.Content が nil でなければ、Content.Type をチェック
		if def.Type == "date" {
			return true
		}
	}
	return false
}
