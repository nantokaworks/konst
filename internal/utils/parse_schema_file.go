package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/nantokaworks/konst/internal/types"
)

// ParseSchemaFile は指定された JSON ファイルまたはディレクトリ内の JSON を再帰的にパースします。
func ParseSchemaFile(filename *string) (*types.Schema, error) {
	info, err := os.Stat(*filename)
	if err != nil {
		return nil, err
	}

	// ディレクトリの場合は再帰的に JSON を処理
	if info.IsDir() {
		master := &types.Schema{
			Definitions: make(map[string]types.Definition),
		}
		err := filepath.Walk(*filename, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") {
				data, err := os.ReadFile(path)
				if err != nil {
					return err
				}
				var schema types.Schema
				if err := json.Unmarshal(data, &schema); err != nil {
					return err
				}
				// master の Version と GoPackage を初回設定
				if master.Version == "" {
					master.Version = schema.Version
				}
				if master.GoPackage == "" {
					master.GoPackage = schema.GoPackage
				}
				for k, def := range schema.Definitions {
					master.Definitions[k] = def // キー重複時は上書き
				}
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
		return master, nil
	}

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
