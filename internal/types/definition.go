package types

// DefinitionType は定義の型を示す列挙型です。
type DefinitionType string

const (
	DefinitionTypeInt   DefinitionType = "int"
	DefinitionTypeInt64 DefinitionType = "int64"
	DefinitionTypeFloat DefinitionType = "float"
	DefinitionTypeDate  DefinitionType = "date"
)

// Definition は各定義の情報を表します。
type Definition struct {
	Type     DefinitionType `json:"type"`
	Value    interface{}    `json:"value"`
	TSMode   TSMode         `json:"tsMode,omitempty"`
	GoMode   GoMode         `json:"goMode,omitempty"`
	DateMode string         `json:"mode,omitempty"`
}
