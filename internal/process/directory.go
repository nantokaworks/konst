package process

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/nantokaworks/konst/internal/i18n"
	"github.com/nantokaworks/konst/internal/types"
	"github.com/nantokaworks/konst/internal/utils"
)

// ProcessDirectory はディレクトリ内のJSONファイルを再帰的に処理します。
func ProcessDirectory(inputDir, outDir string, option *types.CommandOption, isTS bool) error {
	// まず全JSONファイルを読み込んで依存関係を解決
	allDefinitions := make(map[string]types.Definition)
	var jsonFiles []string
	
	err := filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(info.Name(), ".json") {
			return nil
		}
		jsonFiles = append(jsonFiles, path)
		schema, err := utils.ParseSchemaFile(&path)
		if err != nil {
			return err
		}
		// 全定義をマージ
		for name, def := range schema.Definitions {
			allDefinitions[name] = def
		}
		return nil
	})
	if err != nil {
		return err
	}

	// 依存関係を解決
	resolvedDefinitions, err := utils.ResolveDependencies(allDefinitions)
	if err != nil {
		return err
	}

	var tsExports []string
	// 各JSONファイルを処理
	for _, jsonPath := range jsonFiles {
		exportPath, err := ProcessFileWithResolvedDependencies(jsonPath, inputDir, outDir, option, isTS, resolvedDefinitions)
		if err != nil {
			return err
		}
		if isTS && exportPath != "" {
			tsExports = append(tsExports, exportPath)
		}
	}
	// TS出力の場合、index.tsを生成
	if isTS {
		indexPath := filepath.Join(outDir, "index.ts")
		f, err := os.Create(indexPath)
		if err != nil {
			return err
		}
		defer f.Close()
		for _, export := range tsExports {
			line := fmt.Sprintf("export * from './%s';\n", export)
			if _, err := f.WriteString(line); err != nil {
				return err
			}
		}
		fmt.Printf("%s: %s\n", i18n.T(i18n.MsgGenerated), indexPath)
	}
	return nil
}
