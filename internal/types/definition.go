package types

// Definition は各定義の情報を表します。
// Descriptor により "const", "enum", "object" などを指定します。
type Definition struct {
	Descriptor Descriptor        `json:"descriptor"`
	Content    DefinitionContent `json:"content"`
}
