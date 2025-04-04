package types

// Data は全体の JSON 構造です。
type Data struct {
	Version     string                `json:"version"`
	Definitions map[string]Definition `json:"definitions"`
}
