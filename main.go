package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/nantokaworks/konst/internal/process"
	"github.com/nantokaworks/konst/internal/types"
	"github.com/nantokaworks/konst/internal/utils"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] [inputFile]\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
	}
}

// dryRunPreview は生成予定のファイル一覧を表示します
func dryRunPreview(inputPath, outputDir string, option *types.CommandOption) error {
	info, err := os.Stat(inputPath)
	if err != nil {
		return err
	}

	isTS := strings.ToLower(*option.Mode) == "ts"
	ext := ".go"
	if isTS {
		ext = ".ts"
	}

	fmt.Printf("モード: %s\n", *option.Mode)
	fmt.Printf("出力先: %s\n", outputDir)
	fmt.Println("生成予定ファイル:")

	if info.IsDir() {
		var files []string
		err := filepath.Walk(inputPath, func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() || !strings.HasSuffix(info.Name(), ".json") {
				return nil
			}
			rel, err := filepath.Rel(inputPath, path)
			if err != nil {
				return err
			}
			base := strings.TrimSuffix(rel, filepath.Ext(rel))
			outFile := filepath.Join(outputDir, base+ext)
			files = append(files, outFile)
			return nil
		})
		if err != nil {
			return err
		}
		
		// TypeScript の場合はindex.tsも生成される
		if isTS {
			files = append(files, filepath.Join(outputDir, "index.ts"))
		}
		
		for _, file := range files {
			fmt.Printf("  - %s\n", file)
		}
	} else {
		base := strings.TrimSuffix(filepath.Base(inputPath), filepath.Ext(inputPath))
		outFile := filepath.Join(outputDir, base+ext)
		fmt.Printf("  - %s\n", outFile)
	}

	return nil
}

// validateOnly は JSON定義ファイルの検証のみを行います
func validateOnly(inputPath string) error {
	info, err := os.Stat(inputPath)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return filepath.Walk(inputPath, func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() || !strings.HasSuffix(info.Name(), ".json") {
				return nil
			}
			_, err = utils.PaerseSchemaFile(&path)
			if err != nil {
				return fmt.Errorf("ファイル %s: %v", path, err)
			}
			fmt.Printf("✓ %s\n", path)
			return nil
		})
	} else {
		_, err = utils.PaerseSchemaFile(&inputPath)
		if err != nil {
			return err
		}
		fmt.Printf("✓ %s\n", inputPath)
	}
	return nil
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

	// バリデーションモードの場合は検証のみを実行
	if *option.Validate {
		if err := validateOnly(*option.SchemaFile); err != nil {
			fmt.Fprintf(os.Stderr, "バリデーションエラー: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("バリデーション成功: JSON定義に問題ありません")
		return
	}

	// ドライランモードの場合は生成予定ファイル一覧を表示
	if *option.DryRun {
		if err := dryRunPreview(*option.SchemaFile, *option.OutputFile, option); err != nil {
			fmt.Fprintf(os.Stderr, "ドライランエラー: %v\n", err)
			os.Exit(1)
		}
		return
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
