package types

// Schema は全体の JSON 構造です。
type Schema struct {
	Version     string                `json:"version"`
	GoPackage   string                `json:"goPackage"`
	Definitions map[string]Definition `json:"definitions"`
}
