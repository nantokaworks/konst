package types

// DefinitionType は定義の型を示す列挙型です。
type DefinitionType string

const (
	DefinitionTypeInt       DefinitionType = "int"
	DefinitionTypeInt32     DefinitionType = "int32"
	DefinitionTypeInt64     DefinitionType = "int64"
	DefinitionTypeUint      DefinitionType = "uint"
	DefinitionTypeUint32    DefinitionType = "uint32"
	DefinitionTypeUint64    DefinitionType = "uint64"
	DefinitionTypeFloat     DefinitionType = "float" // float = float32
	DefinitionTypeFloat32   DefinitionType = "float32"
	DefinitionTypeFloat64   DefinitionType = "float64"
	DefinitionTypeString    DefinitionType = "string"
	DefinitionTypeBool      DefinitionType = "bool"
	DefinitionTypeDate      DefinitionType = "date"
	DefinitionTypeTimestamp DefinitionType = "timestamp" // 日付のtimestamp型
)

// Definition は各定義の情報を表します。
type Definition struct {
	Type   DefinitionType `json:"type"`
	Value  interface{}    `json:"value"`
	TSMode TSMode         `json:"tsMode,omitempty"`
	GoMode GoMode         `json:"goMode,omitempty"`
	// DateMode フィールドを廃止し、TSModeで統一します。
}
