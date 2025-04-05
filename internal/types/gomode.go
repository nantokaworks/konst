package types

// GoMode は、数字など出力用のモードを示す列挙型です。
type GoMode string

const (
	GoModeInt       GoMode = "int"
	GoModeInt64     GoMode = "int64"
	GoModeTime      GoMode = "time.Time"
	GoModeString    GoMode = "string"
	GoModeTimestamp GoMode = "timestamp"
)
