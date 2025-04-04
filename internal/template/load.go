package template

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func Load(outputFile *string, templateDir *string, spaces *int) (*template.Template, error) {
	// テンプレート関数を定義
	funcMap := createMap(spaces)

	ext := filepath.Ext(*outputFile)
	tmplText, err := loadTemplate(ext, *templateDir)
	if err != nil {
		return nil, err
	}

	tmpl, err := template.New("output").Funcs(funcMap).Parse(tmplText)
	if err != nil {
		return nil, err
	}

	return tmpl.New("output").Funcs(funcMap).Parse(tmplText)
}

// LoadTemplate は、指定されたテンプレートディレクトリから、
// ext に対応するテンプレートファイル (go.tmpl または ts.tmpl) を読み込み、
// 存在しなければ内蔵テンプレートを返します。
func loadTemplate(ext, tmplDir string) (string, error) {
	if tmplDir != "" {
		// ext は ".go" または ".ts" なので、先頭のドットを除いた名前 + ".tmpl"
		templateFile := filepath.Join(tmplDir, strings.TrimPrefix(ext, ".")+".tmpl")
		if _, err := os.Stat(templateFile); err == nil {
			bytes, err := os.ReadFile(templateFile)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		}
	}
	switch ext {
	case ".go":
		return defaultGoTemplate, nil
	case ".ts":
		return defaultTSTemplate, nil
	default:
		return "", fmt.Errorf("未対応の拡張子: %s", ext)
	}
}
