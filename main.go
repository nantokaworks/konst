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

	// 既にディレクトリ入力の場合のチェックはそのまま
	if info.IsDir() {
		if filepath.Ext(*option.OutputFile) != "" {
			fmt.Fprintln(os.Stderr, "エラー: 入力がディレクトリの場合、-o はディレクトリでなければなりません")
			os.Exit(1)
		}
	}

	// 出力モードは、ディレクトリモードの場合、--mode フラグの値で判定
	var isTS bool
	if info.IsDir() {
		if strings.ToLower(*option.Mode) == "ts" {
			isTS = true
		} else {
			isTS = false
		}
	} else {
		// 単一ファイルの場合は outputFile の拡張子から判定
		isTS = strings.HasSuffix(*option.OutputFile, "ts")
	}

	var outDir string
	// 単一ファイルの場合、-oがディレクトリなら入力ファイル名を利用
	if !info.IsDir() && filepath.Ext(*option.OutputFile) == "" {
		outDir = *option.OutputFile
	} else {
		outDir = *option.OutputFile
	}

	if info.IsDir() {
		if err := process.ProcessDirectory(inputPath, outDir, option, isTS); err != nil {
			fmt.Fprintf(os.Stderr, "ディレクトリ処理エラー: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// 単一ファイルの場合
	if filepath.Ext(*option.OutputFile) == "" {
		_, err = process.ProcessFile(inputPath, filepath.Dir(inputPath), outDir, option, isTS)
	} else {
		_, err = process.ProcessFile(inputPath, filepath.Dir(inputPath), outDir, option, isTS)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "ファイル処理エラー: %v\n", err)
		os.Exit(1)
	}
}
