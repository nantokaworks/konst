package process

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/nantokaworks/konst/internal/i18n"
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
	
	// ディレクトリとファイル名を分離
	dir := filepath.Dir(rel)
	fileName := filepath.Base(rel)
	fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))
	
	// ファイル名を命名規則に従って変換
	namingStyle := ""
	if option.NamingStyle != nil {
		namingStyle = *option.NamingStyle
	}
	convertedFileName := utils.ConvertFileName(fileName, namingStyle, isTS)
	
	// ディレクトリも変換
	convertedDir := dir
	if dir != "." {
		convertedDir = utils.ConvertPath(dir, namingStyle, isTS)
	}
	
	// 出力拡張子の決定
	outExt := ".go"
	if isTS {
		outExt = ".ts"
	}
	
	// 出力パスを構築
	var outFilePath string
	
	// Goの場合、goPackageごとにサブディレクトリを作成
	if !isTS && schema.GoPackage != "" {
		// goPackageをディレクトリ名として使用
		packageDir := schema.GoPackage
		if convertedDir != "." {
			outFilePath = filepath.Join(outDir, convertedDir, packageDir, convertedFileName+outExt)
		} else {
			outFilePath = filepath.Join(outDir, packageDir, convertedFileName+outExt)
		}
	} else {
		// TypeScriptまたはgoPackageが空の場合は従来通り
		if convertedDir != "." {
			outFilePath = filepath.Join(outDir, convertedDir, convertedFileName+outExt)
		} else {
			outFilePath = filepath.Join(outDir, convertedFileName+outExt)
		}
	}

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
	fmt.Printf("%s: %s\n", i18n.T(i18n.MsgGenerated), outFilePath)
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
