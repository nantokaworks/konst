package types

// TSMode は、TypeScript 出力で利用するモードを示す列挙型です。
type TSMode string

const (
	TSModeNumber TSMode = "number"
	TSModeBigInt TSMode = "bigint"
	TSModeString TSMode = "string" // 日付を文字列として出力する場合などに利用
	TSModeDate   TSMode = "date"   // 日付を Date 型として出力する場合に利用
)
