package types

// TSMode は、数値および日付出力用のモードを示す列挙型です。
type TSMode string

const (
	ModeNumber    TSMode = "number"
	ModeBigInt    TSMode = "bigint"
	ModeString    TSMode = "string"
	ModeDate      TSMode = "date"
	ModeTimestamp TSMode = "timestamp"
)
