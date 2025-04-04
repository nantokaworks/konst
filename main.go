package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nantokaworks/konst/internal/template"
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

	// コマンドライン引数の解析
	option, err := utils.GetCommandOption()
	if err != nil {
		fmt.Fprintf(os.Stderr, "コマンドライン引数エラー: %v\n", err)
		flag.Usage()
		os.Exit(1)
	}

	// スキーマのパース
	schema, err := utils.PaerseSchemaFile(option.SchemaFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "スキーマファイル読み込みエラー: %v\n", err)
		os.Exit(1)
	}

	// テンプレート読み込み
	tmpl, err := template.Load(option.OutputFile, option.TemplateDir, option.Indent)
	if err != nil {
		fmt.Fprintf(os.Stderr, "テンプレート読み込みエラー: %v\n", err)
		os.Exit(1)
	}

	// 出力ファイルの作成とオープン
	outF := utils.CreateOutputFile(option.OutputFile, option.Force)
	defer outF.Close()

	// テンプレートの実行
	if err := tmpl.Execute(outF, schema); err != nil {
		fmt.Fprintf(os.Stderr, "テンプレート実行エラー: %v\n", err)
		os.Exit(1)
	}

	// 出力完了メッセージ
	fmt.Println("生成完了:", *option.OutputFile)
}
