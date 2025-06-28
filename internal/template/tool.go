package template

import (
	"sort"
	"strings"
	"time"

	"github.com/nantokaworks/konst/internal/types"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// toTitle は文字列をタイトルケースに変換します
func toTitle(s string) string {
	return cases.Title(language.Und, cases.NoLower).String(s)
}

// toCamel は文字列をキャメルケースに変換します（twitch_id -> twitchId）
func toCamel(s string) string {
	if s == "" {
		return s
	}

	// アンダースコアで分割
	parts := strings.Split(s, "_")
	if len(parts) <= 1 {
		return s
	}

	// 最初の部分は小文字のまま、残りは先頭大文字
	result := parts[0]
	for i := 1; i < len(parts); i++ {
		if parts[i] != "" {
			result += cases.Title(language.Und, cases.NoLower).String(parts[i])
		}
	}

	return result
}

// createMap は、テンプレートで使用する関数を定義したマップを作成します。
func sortedKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// indent は、指定されたスペース数のインデントを、文字列の各行の先頭に追加します。
func indent(spaces int, s string) string {
	prefix := strings.Repeat(" ", spaces)
	return prefix + strings.Replace(s, "\n", "\n"+prefix, -1)
}

// tryParseDate は文字列を RFC3339 としてパースできるか試します。
func tryParseDate(s string) (time.Time, bool) {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return time.Time{}, false
	}
	return t, true
}

// hasEnum は定義の中にenum型があるかチェックします
func hasEnum(definitions map[string]types.Definition) bool {
	for _, def := range definitions {
		if def.Type == types.DefinitionTypeEnum {
			return true
		}
	}
	return false
}

// hasTemplate は定義の中にtemplate型があるかチェックします
func hasTemplate(definitions map[string]types.Definition) bool {
	for _, def := range definitions {
		if def.Type == types.DefinitionTypeTemplate {
			return true
		}
	}
	return false
}
