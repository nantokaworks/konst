package process

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/nantokaworks/konst/internal/types"
)

// processDirectory はディレクトリ内のJSONファイルを再帰的に処理します。
func ProcessDirectory(inputDir, outDir string, option *types.CommandOption, isTS bool) error {
	var tsExports []string
	err := filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(info.Name(), ".json") {
			return nil
		}
		exportPath, err := ProcessFile(path, inputDir, outDir, option, isTS)
		if err != nil {
			return err
		}
		if isTS && exportPath != "" {
			tsExports = append(tsExports, exportPath)
		}
		return nil
	})
	if err != nil {
		return err
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
		fmt.Println("生成完了:", indexPath)
	}
	return nil
}
