package types

// GoMode は、数字など出力用のモードを示す列挙型です。
type GoMode string

const (
	GoModeNumber GoMode = "number"
	GoModeBigInt GoMode = "bigint"
)
