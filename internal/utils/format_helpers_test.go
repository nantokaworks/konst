package utils

import (
	"testing"

	"github.com/nantokaworks/konst/internal/types"
)

func TestAsString(t *testing.T) {
	tests := []struct {
		input    any
		expected string
	}{
		{"hello", "hello"},
		{42, "42"},
		{true, "true"},
		{3.14, "3.14"},
		{nil, "<nil>"},
	}

	for _, tt := range tests {
		result := AsString(tt.input)
		if result != tt.expected {
			t.Errorf("AsString(%v) = %q, expected %q", tt.input, result, tt.expected)
		}
	}
}

func TestMonthConst(t *testing.T) {
	tests := []struct {
		month    int
		expected string
	}{
		{1, "time.January"},
		{2, "time.February"},
		{6, "time.June"},
		{12, "time.December"},
		{0, "time.January"},  // boundary case
		{13, "time.January"}, // boundary case
		{-1, "time.January"}, // boundary case
	}

	for _, tt := range tests {
		result := MonthConst(tt.month)
		if result != tt.expected {
			t.Errorf("MonthConst(%d) = %q, expected %q", tt.month, result, tt.expected)
		}
	}
}

func TestHasDate(t *testing.T) {
	tests := []struct {
		name        string
		definitions map[string]types.Definition
		expected    bool
	}{
		{
			name: "Has date",
			definitions: map[string]types.Definition{
				"CreatedAt": {Type: types.DefinitionTypeDate},
				"Name":      {Type: types.DefinitionTypeString},
			},
			expected: true,
		},
		{
			name: "No date",
			definitions: map[string]types.Definition{
				"Name": {Type: types.DefinitionTypeString},
				"Age":  {Type: types.DefinitionTypeInt},
			},
			expected: false,
		},
		{
			name:        "Empty definitions",
			definitions: map[string]types.Definition{},
			expected:    false,
		},
		{
			name:        "Nil definitions",
			definitions: nil,
			expected:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HasDate(tt.definitions)
			if result != tt.expected {
				t.Errorf("HasDate() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestConvertTSType(t *testing.T) {
	tests := []struct {
		goType   string
		expected string
	}{
		{"int", "number"},
		{"int32", "number"},
		{"int64", "number"},
		{"uint", "number"},
		{"float32", "number"},
		{"float64", "number"},
		{"string", "string"},
		{"bool", "boolean"},
		{"date", "Date"},
		{"timestamp", "number"},
		{"int[]", "number[]"},
		{"string[]", "string[]"},
		{"bool[]", "boolean[]"},
		{"date[]", "Date[]"},
		{"unknown", "any"},
		{"custom", "any"},
	}

	for _, tt := range tests {
		result := ConvertTSType(tt.goType)
		if result != tt.expected {
			t.Errorf("ConvertTSType(%q) = %q, expected %q", tt.goType, result, tt.expected)
		}
	}
}

func TestConvertTSTypeNestedArrays(t *testing.T) {
	// ネストした配列型のテスト
	tests := []struct {
		goType   string
		expected string
	}{
		{"int[][]", "number[][]"},
		{"string[][][]", "string[][][]"},
	}

	for _, tt := range tests {
		result := ConvertTSType(tt.goType)
		if result != tt.expected {
			t.Errorf("ConvertTSType(%q) = %q, expected %q", tt.goType, result, tt.expected)
		}
	}
}
