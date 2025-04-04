package types

type ConstContent struct {
	Type   string      `json:"type"`
	Value  interface{} `json:"value"`
	TSMode TSMode      `json:"tsMode,omitempty"` // 変更：string -> TSMode
}
