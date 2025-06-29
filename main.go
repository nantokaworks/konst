package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/nantokaworks/konst/internal/i18n"
	"github.com/nantokaworks/konst/internal/process"
	"github.com/nantokaworks/konst/internal/types"
	"github.com/nantokaworks/konst/internal/utils"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] <inputDirectory>\n", os.Args[0])
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

	fmt.Printf("%s: %s\n", i18n.T(i18n.MsgMode), *option.Mode)
	fmt.Printf("%s: %s\n", i18n.T(i18n.MsgOutputDirectory), outputDir)
	fmt.Printf("%s:\n", i18n.T(i18n.MsgFilesToBeGenerated))

	if !info.IsDir() {
		return fmt.Errorf(i18n.T(i18n.MsgInputMustBeDir))
	}

	var files []string
	err = filepath.Walk(inputPath, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(info.Name(), ".json") {
			return nil
		}
		rel, err := filepath.Rel(inputPath, path)
		if err != nil {
			return err
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
		
		// 出力パスを構築
		var outFile string
		if convertedDir != "." {
			outFile = filepath.Join(outputDir, convertedDir, convertedFileName+ext)
		} else {
			outFile = filepath.Join(outputDir, convertedFileName+ext)
		}
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

	return nil
}

// validateOnly は JSON定義ファイルの検証のみを行います
func validateOnly(inputPath string) error {
	info, err := os.Stat(inputPath)
	if err != nil {
		return err
	}

	if !info.IsDir() {
		return fmt.Errorf(i18n.T(i18n.MsgInputMustBeDir))
	}

	return filepath.Walk(inputPath, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(info.Name(), ".json") {
			return nil
		}
		_, err = utils.ParseSchemaFile(&path)
		if err != nil {
			return fmt.Errorf("%s %s: %v", i18n.T(i18n.MsgFileError), path, err)
		}
		fmt.Printf("✓ %s\n", path)
		return nil
	})
}

func main() {
	option, err := utils.GetCommandOption()
	if err != nil {
		// i18n初期化前なのでデフォルトメッセージを使用
		fmt.Fprintf(os.Stderr, "Command line argument error: %v\n", err)
		flag.Usage()
		os.Exit(1)
	}

	// i18nシステムを初期化
	if err := i18n.Init(*option.Locale); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize i18n: %v\n", err)
		// エラーでも続行（英語デフォルトで動作）
	}

	// バリデーションモードの場合は検証のみを実行
	if *option.Validate {
		if err := validateOnly(*option.SchemaFile); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", i18n.T(i18n.MsgValidationError), err)
			os.Exit(1)
		}
		fmt.Println(i18n.T(i18n.MsgValidationSuccess))
		return
	}

	// ドライランモードの場合は生成予定ファイル一覧を表示
	if *option.DryRun {
		if err := dryRunPreview(*option.SchemaFile, *option.OutputFile, option); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", i18n.T(i18n.MsgDryRunError), err)
			os.Exit(1)
		}
		return
	}

	inputPath := *option.SchemaFile
	info, err := os.Stat(inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", i18n.T(i18n.MsgInputPathError), err)
		os.Exit(1)
	}

	// ディレクトリのみ対応
	if !info.IsDir() {
		fmt.Fprintln(os.Stderr, i18n.T(i18n.MsgInputMustBeDir))
		os.Exit(1)
	}

	// 出力先が拡張子付きファイル名の場合はエラー
	if filepath.Ext(*option.OutputFile) != "" {
		fmt.Fprintln(os.Stderr, i18n.T(i18n.MsgOutputMustBeDir))
		os.Exit(1)
	}

	// 出力モードは --mode フラグで判定
	isTS := strings.ToLower(*option.Mode) == "ts"

	if err := process.ProcessDirectory(inputPath, *option.OutputFile, option, isTS); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", i18n.T(i18n.MsgProcessingError), err)
		os.Exit(1)
	}
}
