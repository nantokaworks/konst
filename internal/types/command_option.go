package types

type CommandOption struct {
	SchemaFile   *string // スキーマファイル名
	OutputFile   *string // 出力ファイル名
	TemplateDir  *string // テンプレートディレクトリ
	Force        *bool   // 強制オプション
	Indent       *int    // インデント数
	Mode         *string // 出力モード (go, ts)
	Validate     *bool   // バリデーションのみ
	DryRun       *bool   // ドライラン
	Watch        *bool   // ウォッチモード
	NamingStyle  *string // ファイル命名規則 (kebab, camel, snake)
	Locale       *string // 言語設定 (ja, en)
}
