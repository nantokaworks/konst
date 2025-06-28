package template

import (
	"strings"
	"text/template"
)

func createMap(spaces *int) map[string]interface{} {

	indentLevel := func(level int, s string) string {
		return indent(level*(*spaces), s)
	}

	return template.FuncMap{
		"formatGo":         formatGo,
		"formatTS":         formatTS,
		"formatConstValue": formatConstValue,
		"convertTSType":    convertTSType,
		"indent":           indentLevel,
		"sortedKeys":       sortedKeys,
		"title":            title, // 変更: strings.Title から独自の title 関数へ
		"toTitle":          toTitle, // 追加: toTitle関数
		"asString":         asString,
		"hasDate":          hasDate,
		"hasEnum":          hasEnum, // 追加: hasEnum関数
		"contains":         strings.Contains, // 追加: contains関数
	}
}
