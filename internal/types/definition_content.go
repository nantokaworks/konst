package types

import (
	"encoding/json"
)

// EnumContent は "enum" 定義用です。
type EnumContent struct {
	Values map[string]string `json:"values"`
}

// FieldDefinition は object の各フィールド定義です。
type FieldDefinition struct {
	Type string `json:"type"`
}

// ObjectContent は "object" 定義用です。
type ObjectContent struct {
	Fields map[string]FieldDefinition `json:"fields"`
}

// DefinitionContent は、定義内容のユニオンです。
// UnmarshalJSON で、内容にあわせて ConstContent, EnumContent または ObjectContent を設定します。
type DefinitionContent struct {
	// Only one will be non-nil.
	ConstContent  *ConstContent
	EnumContent   *EnumContent
	ObjectContent *ObjectContent
}

func (d *DefinitionContent) UnmarshalJSON(data []byte) error {
	var temp map[string]json.RawMessage
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	// 判定キーでコンテンツの種類を決定
	if _, ok := temp["values"]; ok {
		var ec EnumContent
		if err := json.Unmarshal(data, &ec); err != nil {
			return err
		}
		d.EnumContent = &ec
		return nil
	}
	if _, ok := temp["fields"]; ok {
		var oc ObjectContent
		if err := json.Unmarshal(data, &oc); err != nil {
			return err
		}
		d.ObjectContent = &oc
		return nil
	}
	// それ以外は定数用とする
	var cc ConstContent
	if err := json.Unmarshal(data, &cc); err != nil {
		return err
	}
	d.ConstContent = &cc
	return nil
}
