package template

import (
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
		"asString":         asString,
		"hasDate":          hasDate,
	}
}
