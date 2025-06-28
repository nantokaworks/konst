package template

import (
	"sort"
	"strings"
	"time"

	"github.com/nantokaworks/konst/internal/types"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// title は文字列の先頭を大文字に変換します（strings.Title の代替）。
func title(s string) string {
	return cases.Title(language.Und, cases.NoLower).String(s)
}

// toTitle は文字列をタイトルケースに変換します
func toTitle(s string) string {
	return cases.Title(language.Und, cases.NoLower).String(s)
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
