package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/nantokaworks/konst/internal/process"
	"github.com/nantokaworks/konst/internal/utils"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] [inputFile]\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
	}
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] [inputFile]\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
	}

	option, err := utils.GetCommandOption()
	if err != nil {
		fmt.Fprintf(os.Stderr, "コマンドライン引数エラー: %v\n", err)
		flag.Usage()
		os.Exit(1)
	}

	inputPath := *option.SchemaFile
	info, err := os.Stat(inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "入力パス取得エラー: %v\n", err)
		os.Exit(1)
	}

	// 出力先が拡張子付きファイル名の場合はエラー
	if filepath.Ext(*option.OutputFile) != "" {
		fmt.Fprintln(os.Stderr, "エラー: -o にはディレクトリを指定してください")
		os.Exit(1)
	}

	// 出力モードは --mode フラグで判定
	isTS := strings.ToLower(*option.Mode) == "ts"

	// 入力がファイルの場合、親ディレクトリを入力ディレクトリとして処理
	var inputDir string
	if info.IsDir() {
		inputDir = inputPath
	} else {
		inputDir = filepath.Dir(inputPath)
	}

	if err := process.ProcessDirectory(inputDir, *option.OutputFile, option, isTS); err != nil {
		fmt.Fprintf(os.Stderr, "処理エラー: %v\n", err)
		os.Exit(1)
	}
}
