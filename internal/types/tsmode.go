package types

// TSMode は、数字など出力用のモードを示す列挙型です。
type TSMode string

const (
	ModeNumber TSMode = "number"
	ModeBigInt TSMode = "bigint"
)
