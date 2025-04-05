package types

// DateMode は、日付出力用のモードを示す列挙型です。
type DateMode string

const (
	DateModeString DateMode = "string"
	DateModeDate   DateMode = "date"
)
