package utils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/nantokaworks/konst/internal/types"
)

func TestParseSchemaFile(t *testing.T) {
	// テスト用の一時ディレクトリ作成
	tempDir := t.TempDir()

	// テスト用JSONファイル作成
	testJSON := `{
		"version": "1.0",
		"goPackage": "test",
		"definitions": {
			"TestString": {
				"type": "string",
				"value": "hello"
			},
			"TestInt": {
				"type": "int",
				"value": 42
			}
		}
	}`

	testFile := filepath.Join(tempDir, "test.json")
	err := os.WriteFile(testFile, []byte(testJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// テスト実行
	schema, err := ParseSchemaFile(&testFile)
	if err != nil {
		t.Fatalf("ParseSchemaFile failed: %v", err)
	}

	// 検証
	if schema.Version != "1.0" {
		t.Errorf("Expected version '1.0', got '%s'", schema.Version)
	}

	if schema.GoPackage != "test" {
		t.Errorf("Expected goPackage 'test', got '%s'", schema.GoPackage)
	}

	if len(schema.Definitions) != 2 {
		t.Errorf("Expected 2 definitions, got %d", len(schema.Definitions))
	}

	// 個別定義の検証
	testString, exists := schema.Definitions["TestString"]
	if !exists {
		t.Error("TestString definition not found")
	}
	if testString.Type != types.DefinitionTypeString {
		t.Errorf("Expected string type, got %s", testString.Type)
	}
	if testString.Value != "hello" {
		t.Errorf("Expected value 'hello', got %v", testString.Value)
	}
}

func TestParseSchemaFileDirectory(t *testing.T) {
	// テスト用の一時ディレクトリ作成
	tempDir := t.TempDir()

	// 複数のJSONファイル作成
	json1 := `{
		"version": "1.0",
		"goPackage": "test",
		"definitions": {
			"Value1": {
				"type": "string",
				"value": "test1"
			}
		}
	}`

	json2 := `{
		"version": "1.0",
		"definitions": {
			"Value2": {
				"type": "int",
				"value": 123
			}
		}
	}`

	file1 := filepath.Join(tempDir, "test1.json")
	file2 := filepath.Join(tempDir, "test2.json")

	err := os.WriteFile(file1, []byte(json1), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file 1: %v", err)
	}

	err = os.WriteFile(file2, []byte(json2), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file 2: %v", err)
	}

	// ディレクトリ指定でテスト実行
	schema, err := ParseSchemaFile(&tempDir)
	if err != nil {
		t.Fatalf("ParseSchemaFile failed: %v", err)
	}

	// 検証
	if len(schema.Definitions) != 2 {
		t.Errorf("Expected 2 definitions, got %d", len(schema.Definitions))
	}

	// 両方の定義が存在することを確認
	_, exists1 := schema.Definitions["Value1"]
	_, exists2 := schema.Definitions["Value2"]

	if !exists1 {
		t.Error("Value1 definition not found")
	}
	if !exists2 {
		t.Error("Value2 definition not found")
	}
}

func TestParseSchemaFileInvalidJSON(t *testing.T) {
	tempDir := t.TempDir()

	// 不正なJSONファイル作成
	invalidJSON := `{
		"version": "1.0"
		"goPackage": "test"  // missing comma
	}`

	testFile := filepath.Join(tempDir, "invalid.json")
	err := os.WriteFile(testFile, []byte(invalidJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// エラーが発生することを確認
	_, err = ParseSchemaFile(&testFile)
	if err == nil {
		t.Error("Expected error for invalid JSON, but got nil")
	}
}
