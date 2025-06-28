package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateOutputFile(outputFile *string, force *bool) *os.File {

	// 書き出し先のディレクトリが存在しない場合、自動で作成する
	outDir := filepath.Dir(*outputFile)
	if err := os.MkdirAll(outDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "出力先ディレクトリ作成エラー: %v\n", err)
		os.Exit(1)
	}

	// 出力ファイルが既に存在するかチェックする
	if _, err := os.Stat(*outputFile); err == nil {
		if !*force {
			fmt.Fprintf(os.Stderr, "警告: 出力ファイル %s は既に存在します。-f オプションを指定して強制上書きしてください。\n", *outputFile)
			os.Exit(1)
		}
	}

	outF, err := os.Create(*outputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "出力ファイル作成エラー: %v\n", err)
		os.Exit(1)
	}

	return outF
}
