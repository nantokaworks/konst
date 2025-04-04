package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/nantokaworks/konst/internal/types"
)

func PaerseSchemaFile(filePath *string) (*types.Data, error) {

	jsonBytes, err := os.ReadFile(*filePath)
	if err != nil {
		return nil, fmt.Errorf("JSONファイル読み込みエラー: %v", err)
	}

	var data types.Data
	if err := json.Unmarshal(jsonBytes, &data); err != nil {
		return nil, fmt.Errorf("JSONパースエラー: %v", err)
	}

	return &data, nil
}
