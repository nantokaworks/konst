package utils

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/nantokaworks/konst/internal/i18n"
	"github.com/nantokaworks/konst/internal/types"
)

const VERSION string = "v0.3.6"

func GetCommandOption() (*types.CommandOption, error) {

	// --localeフラグを事前にチェック（ヘルプメッセージのため）
	locale := ""
	for i, arg := range os.Args {
		if arg == "--locale" && i+1 < len(os.Args) {
			locale = os.Args[i+1]
			break
		} else if len(arg) > 9 && arg[:9] == "--locale=" {
			locale = arg[9:]
			break
		}
	}
	
	// ヘルプメッセージの初期化（flag.Parse前に呼ぶ必要がある）
	i18n.InitHelpMessagesWithLocale(locale)

	// コマンドライン引数のパース
	schemaFile := flag.String("i", "", i18n.GetHelpMessage(i18n.HelpInputFile))
	outputFile := flag.String("o", "", i18n.GetHelpMessage(i18n.HelpOutputDir))
	templateDirFlag := flag.String("t", "", i18n.GetHelpMessage(i18n.HelpTemplateDir))
	forceFlag := flag.Bool("f", false, i18n.GetHelpMessage(i18n.HelpForce))
	indentFlag := flag.Int("indent", 2, i18n.GetHelpMessage(i18n.HelpIndent))
	versionFlag := flag.Bool("v", false, i18n.GetHelpMessage(i18n.HelpVersion))
	versionLFlag := flag.Bool("version", false, i18n.GetHelpMessage(i18n.HelpVersion))
	modeFlag := flag.String("m", "go", i18n.GetHelpMessage(i18n.HelpMode))
	validateFlag := flag.Bool("validate", false, i18n.GetHelpMessage(i18n.HelpValidate))
	dryRunFlag := flag.Bool("dry-run", false, i18n.GetHelpMessage(i18n.HelpDryRun))
	watchFlag := flag.Bool("watch", false, i18n.GetHelpMessage(i18n.HelpWatch))
	namingStyleFlag := flag.String("naming", "", i18n.GetHelpMessage(i18n.HelpNaming))
	localeFlag := flag.String("locale", "", i18n.GetHelpMessage(i18n.HelpLocale))
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
		return nil, fmt.Errorf("please specify output filename with -o option")
	}

	// テンプレートディレクトリのパスを取得
	tmplDir := *templateDirFlag
	if tmplDir == "" {
		tmplDir = os.Getenv("KONST_TEMPLATES")
	}
	if tmplDir == "" {
		exePath, err := os.Executable()
		if err != nil {
			return nil, fmt.Errorf("executable path error: %v", err)
		}
		tmplDir = filepath.Join(filepath.Dir(exePath), "templates")
	}

	// 言語設定を決定（優先順位順）
	finalLocale := *localeFlag
	if finalLocale == "" {
		// 1. KONST_LOCALE環境変数をチェック
		finalLocale = os.Getenv("KONST_LOCALE")
	}
	if finalLocale == "" {
		// 2. システムロケールを自動検出
		finalLocale = i18n.DetectSystemLocale()
	}
	if finalLocale == "" {
		// 3. 最終的なフォールバック
		finalLocale = "en"
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
		NamingStyle: namingStyleFlag,
		Locale:      &finalLocale,
	}, nil
}
