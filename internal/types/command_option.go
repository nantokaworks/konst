package types

type CommandOption struct {
	SchemaFile  *string // スキーマファイル名
	OutputFile  *string // 出力ファイル名
	TemplateDir *string // テンプレートディレクトリ
	Force       *bool   // 強制オプション
	Indent      *int    // インデント数
}
