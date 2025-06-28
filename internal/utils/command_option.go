package utils

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/nantokaworks/konst/internal/types"
)

const VERSION string = "v0.3.0"

func GetCommandOption() (*types.CommandOption, error) {

	// コマンドライン引数のパース
	schemaFile := flag.String("i", "", "定数定義のJSONファイルまたはディレクトリ（指定がなければ最初の引数を使用）")
	outputFile := flag.String("o", "", "出力先ディレクトリ（必須）")
	templateDirFlag := flag.String("t", "", "カスタムテンプレートディレクトリのパス（省略時は環境変数 KONST_TEMPLATES、なければ実行ファイルと同じ場所のtemplatesディレクトリを使用）")
	forceFlag := flag.Bool("f", false, "既存ファイルを強制的に上書きする")
	indentFlag := flag.Int("indent", 2, "インデント数（デフォルトは2）")
	versionFlag := flag.Bool("v", false, "バージョンを表示する")
	versionLFlag := flag.Bool("version", false, "バージョンを表示する")
	modeFlag := flag.String("m", "go", "出力モードを指定する（go, ts）")
	validateFlag := flag.Bool("validate", false, "JSON定義の検証のみを行う（コード生成は行わない）")
	dryRunFlag := flag.Bool("dry-run", false, "実際の生成は行わず、生成予定のファイル一覧を表示する")
	watchFlag := flag.Bool("watch", false, "ファイル変更を監視して自動生成する（実験的機能）")
	flag.Parse()

	// バージョン表示処理
	if *versionFlag || *versionLFlag {
		fmt.Printf("Konst version %s\n", VERSION)
		os.Exit(0)
	}

	//  スキーマファイルのパスを取得
	inFile := *schemaFile
	if inFile == "" {
		if flag.NArg() > 0 {
			inFile = flag.Arg(0)
		} else {
			inFile = "konst.json"
		}
	}

	// 出力先のチェック（バリデーションモード以外では必須）
	if *outputFile == "" && !*validateFlag {
		return nil, fmt.Errorf("出力ファイル名を -o オプションで指定してください")
	}

	// テンプレートディレクトリのパスを取得
	tmplDir := *templateDirFlag
	if tmplDir == "" {
		tmplDir = os.Getenv("KONST_TEMPLATES")
	}
	if tmplDir == "" {
		exePath, err := os.Executable()
		if err != nil {
			return nil, fmt.Errorf("実行ファイルパス取得エラー: %v", err)
		}
		tmplDir = filepath.Join(filepath.Dir(exePath), "templates")
	}

	return &types.CommandOption{
		SchemaFile:  &inFile,
		OutputFile:  outputFile,
		TemplateDir: &tmplDir,
		Force:       forceFlag,
		Indent:      indentFlag,
		Mode:        modeFlag,
		Validate:    validateFlag,
		DryRun:      dryRunFlag,
		Watch:       watchFlag,
	}, nil
}
