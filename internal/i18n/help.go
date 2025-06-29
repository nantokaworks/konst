package i18n

import "os"

// ヘルプメッセージのキー定義
const (
	HelpInputFile      = "help_input_file"
	HelpOutputDir      = "help_output_dir"
	HelpTemplateDir    = "help_template_dir"
	HelpForce          = "help_force"
	HelpIndent         = "help_indent"
	HelpVersion        = "help_version"
	HelpMode           = "help_mode"
	HelpValidate       = "help_validate"
	HelpDryRun         = "help_dry_run"
	HelpWatch          = "help_watch"
	HelpNaming         = "help_naming"
	HelpLocale         = "help_locale"
)

// helpLocale はヘルプメッセージ用のロケール設定を保持
var helpLocale string

// GetHelpMessage は現在の言語設定に基づいてヘルプメッセージを返します
func GetHelpMessage(key string) string {
	// デフォルト（英語）のヘルプメッセージ
	defaultHelp := map[string]string{
		HelpInputFile:   "JSON file or directory for constant definitions (uses first argument if not specified)",
		HelpOutputDir:   "Output directory (required)",
		HelpTemplateDir: "Custom template directory path (uses KONST_TEMPLATES env var if omitted, or templates directory in same location as executable)",
		HelpForce:       "Force overwrite existing files",
		HelpIndent:      "Number of indents (default is 2)",
		HelpVersion:     "Show version",
		HelpMode:        "Specify output mode (go, ts)",
		HelpValidate:    "Only validate JSON definitions (no code generation)",
		HelpDryRun:      "Show list of files to be generated without actual generation",
		HelpWatch:       "Monitor file changes for automatic generation (experimental feature)",
		HelpNaming:      "File naming convention (kebab, camel, snake) - TypeScript defaults to kebab, Go defaults to snake",
		HelpLocale:      "Language setting (ja, en) - uses KONST_LOCALE env var if not specified, then auto-detects system locale",
	}

	// 日本語のヘルプメッセージ
	japaneseHelp := map[string]string{
		HelpInputFile:   "定数定義のJSONファイルまたはディレクトリ（指定がなければ最初の引数を使用）",
		HelpOutputDir:   "出力先ディレクトリ（必須）",
		HelpTemplateDir: "カスタムテンプレートディレクトリのパス（省略時は環境変数 KONST_TEMPLATES、なければ実行ファイルと同じ場所のtemplatesディレクトリを使用）",
		HelpForce:       "既存ファイルを強制的に上書きする",
		HelpIndent:      "インデント数（デフォルトは2）",
		HelpVersion:     "バージョンを表示する",
		HelpMode:        "出力モードを指定する（go, ts）",
		HelpValidate:    "JSON定義の検証のみを行う（コード生成は行わない）",
		HelpDryRun:      "実際の生成は行わず、生成予定のファイル一覧を表示する",
		HelpWatch:       "ファイル変更を監視して自動生成する（実験的機能）",
		HelpNaming:      "ファイル命名規則（kebab, camel, snake）TypeScriptはデフォルトでkebab、Goはデフォルトでsnake",
		HelpLocale:      "言語設定（ja, en）未指定時は環境変数KONST_LOCALE、次にシステムロケールを自動検出",
	}

	// 初期化時に設定されたロケールを使用
	locale := helpLocale
	if locale == "" {
		locale = "en" // デフォルト
	}
	
	// 言語に応じたメッセージを返す
	if locale == "ja" {
		if msg, ok := japaneseHelp[key]; ok {
			return msg
		}
	}
	
	// デフォルト（英語）を返す
	if msg, ok := defaultHelp[key]; ok {
		return msg
	}
	
	return key
}

// InitHelpMessages はヘルプメッセージを初期化します
// flag.Parse()の前に呼ぶ必要があります
func InitHelpMessages() {
	InitHelpMessagesWithLocale("")
}

// InitHelpMessagesWithLocale は指定されたロケールでヘルプメッセージを初期化します
func InitHelpMessagesWithLocale(cmdLineLocale string) {
	locale := cmdLineLocale
	
	// コマンドラインで指定されていない場合、環境変数とシステムロケールから判定
	if locale == "" {
		// 環境変数をチェック
		if envLocale := os.Getenv("KONST_LOCALE"); envLocale != "" {
			locale = envLocale
		} else {
			// システムロケールを自動検出
			locale = DetectSystemLocale()
		}
	}
	
	// デフォルト値を設定
	if locale == "" {
		locale = "en"
	}
	
	// グローバル変数に保存してGetHelpMessageで使用
	helpLocale = locale
}