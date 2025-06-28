package template

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/nantokaworks/konst/internal/utils"
)

func createMap(spaces *int) map[string]interface{} {

	indentLevel := func(level int, s string) string {
		return indent(level*(*spaces), s)
	}

	return template.FuncMap{
		"formatGo":         formatGo,
		"formatTS":         formatTS,
		"formatConstValue": formatConstValue,
		"convertTSType":    utils.ConvertTSType,
		"indent":           indentLevel,
		"sortedKeys":       sortedKeys,
		"title":            toTitle, // strings.Title から toTitle 関数へ
		"toTitle":          toTitle, // 追加: toTitle関数
		"toCamel":          toCamel, // 追加: toCamel関数
		"asString":         utils.AsString,
		"hasDate":          utils.HasDate,
		"hasEnum":          hasEnum,          // 追加: hasEnum関数
		"hasTemplate":      hasTemplate,      // 追加: hasTemplate関数
		"contains":         strings.Contains, // 追加: contains関数
		"printf":           fmt.Sprintf,      // 追加: printf関数
	}
}
