package utils

import (
	"encoding/json"
	"os"

	"github.com/nantokaworks/konst/internal/types"
)

// PaerseSchemaFile は指定された JSON ファイルをパースします。
func PaerseSchemaFile(filename *string) (*types.Schema, error) {
	data, err := os.ReadFile(*filename)
	if err != nil {
		return nil, err
	}
	var schema types.Schema
	if err := json.Unmarshal(data, &schema); err != nil {
		return nil, err
	}

	return &schema, nil
}
