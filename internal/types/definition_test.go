package types

import (
	"encoding/json"
	"testing"
)

func TestDefinitionTypeMarshal(t *testing.T) {
	tests := []struct {
		name     string
		defType  DefinitionType
		expected string
	}{
		{"String type", DefinitionTypeString, `"string"`},
		{"Int type", DefinitionTypeInt, `"int"`},
		{"Bool type", DefinitionTypeBool, `"bool"`},
		{"Enum type", DefinitionTypeEnum, `"enum"`},
		{"Template type", DefinitionTypeTemplate, `"template"`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.defType)
			if err != nil {
				t.Fatalf("Marshal failed: %v", err)
			}

			if string(data) != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, string(data))
			}
		})
	}
}

func TestDefinitionUnmarshal(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		expected Definition
	}{
		{
			name: "String definition",
			jsonData: `{
				"type": "string",
				"value": "hello"
			}`,
			expected: Definition{
				Type:  DefinitionTypeString,
				Value: "hello",
			},
		},
		{
			name: "Enum definition",
			jsonData: `{
				"type": "enum",
				"values": ["active", "inactive"],
				"default": "active"
			}`,
			expected: Definition{
				Type:    DefinitionTypeEnum,
				Values:  []string{"active", "inactive"},
				Default: "active",
			},
		},
		{
			name: "Template definition",
			jsonData: `{
				"type": "template",
				"template": "user:%id%",
				"parameters": ["id"]
			}`,
			expected: Definition{
				Type:       DefinitionTypeTemplate,
				Template:   "user:%id%",
				Parameters: []string{"id"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var def Definition
			err := json.Unmarshal([]byte(tt.jsonData), &def)
			if err != nil {
				t.Fatalf("Unmarshal failed: %v", err)
			}

			if def.Type != tt.expected.Type {
				t.Errorf("Expected type %s, got %s", tt.expected.Type, def.Type)
			}

			if def.Value != tt.expected.Value {
				t.Errorf("Expected value %v, got %v", tt.expected.Value, def.Value)
			}

			if len(def.Values) != len(tt.expected.Values) {
				t.Errorf("Expected %d values, got %d", len(tt.expected.Values), len(def.Values))
			}

			if def.Default != tt.expected.Default {
				t.Errorf("Expected default %s, got %s", tt.expected.Default, def.Default)
			}

			if def.Template != tt.expected.Template {
				t.Errorf("Expected template %s, got %s", tt.expected.Template, def.Template)
			}

			if len(def.Parameters) != len(tt.expected.Parameters) {
				t.Errorf("Expected %d parameters, got %d", len(tt.expected.Parameters), len(def.Parameters))
			}
		})
	}
}

func TestSchemaUnmarshal(t *testing.T) {
	jsonData := `{
		"version": "1.0",
		"goPackage": "test",
		"definitions": {
			"TestString": {
				"type": "string",
				"value": "hello"
			},
			"TestEnum": {
				"type": "enum",
				"values": ["a", "b", "c"]
			}
		}
	}`

	var schema Schema
	err := json.Unmarshal([]byte(jsonData), &schema)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if schema.Version != "1.0" {
		t.Errorf("Expected version '1.0', got '%s'", schema.Version)
	}

	if schema.GoPackage != "test" {
		t.Errorf("Expected goPackage 'test', got '%s'", schema.GoPackage)
	}

	if len(schema.Definitions) != 2 {
		t.Errorf("Expected 2 definitions, got %d", len(schema.Definitions))
	}

	testString, exists := schema.Definitions["TestString"]
	if !exists {
		t.Error("TestString definition not found")
	}
	if testString.Type != DefinitionTypeString {
		t.Errorf("Expected string type, got %s", testString.Type)
	}

	testEnum, exists := schema.Definitions["TestEnum"]
	if !exists {
		t.Error("TestEnum definition not found")
	}
	if testEnum.Type != DefinitionTypeEnum {
		t.Errorf("Expected enum type, got %s", testEnum.Type)
	}
}
