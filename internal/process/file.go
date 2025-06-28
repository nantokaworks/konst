package process

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/nantokaworks/konst/internal/template"
	"github.com/nantokaworks/konst/internal/types"
	"github.com/nantokaworks/konst/internal/utils"
)

// processFile は1つのJSONファイルを処理します。
func ProcessFile(jsonPath, inputDir, outDir string, option *types.CommandOption, isTS bool) (string, error) {
	// 個々のJSONファイルをパース
	schema, err := utils.ParseSchemaFile(&jsonPath)
	if err != nil {
		return "", err
	}
	// 入力ディレクトリからの相対パス取得
	rel, err := filepath.Rel(inputDir, jsonPath)
	if err != nil {
		return "", err
	}
	base := strings.TrimSuffix(rel, filepath.Ext(rel))
	// 出力拡張子の決定
	outExt := ".go"
	if isTS {
		outExt = ".ts"
	}
	outFilePath := filepath.Join(outDir, base+outExt)

	tmpl, err := template.Load(&outFilePath, option.TemplateDir, option.Indent)
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(filepath.Dir(outFilePath), 0755); err != nil {
		return "", err
	}
	outF := utils.CreateOutputFile(&outFilePath, option.Force)
	defer outF.Close()
	if err := tmpl.Execute(outF, schema); err != nil {
		return "", err
	}
	fmt.Println("生成完了:", outFilePath)
	// TS出力の場合、相対パスを返す（拡張子抜き）
	if isTS {
		relOut, err := filepath.Rel(outDir, outFilePath)
		if err != nil {
			return "", err
		}
		return strings.TrimSuffix(filepath.ToSlash(relOut), ".ts"), nil
	}
	return "", nil
}
