package template

import (
	"testing"

	"github.com/nantokaworks/konst/internal/types"
)

func TestToTitle(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "Hello"},
		{"world", "World"},
		{"", ""},
		{"a", "A"},
		{"TEST", "TEST"},
		{"test_case", "Test_case"},
	}

	for _, tt := range tests {
		result := toTitle(tt.input)
		if result != tt.expected {
			t.Errorf("toTitle(%q) = %q, expected %q", tt.input, result, tt.expected)
		}
	}
}

func TestToCamel(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"twitch_id", "twitchId"},
		{"user_name", "userName"},
		{"api_key", "apiKey"},
		{"single", "single"},
		{"", ""},
		{"_test", "Test"},
		{"test_", "test"},
		{"multiple_under_scores", "multipleUnderScores"},
	}

	for _, tt := range tests {
		result := toCamel(tt.input)
		if result != tt.expected {
			t.Errorf("toCamel(%q) = %q, expected %q", tt.input, result, tt.expected)
		}
	}
}

func TestHasEnum(t *testing.T) {
	tests := []struct {
		name        string
		definitions map[string]types.Definition
		expected    bool
	}{
		{
			name: "Has enum",
			definitions: map[string]types.Definition{
				"Status": {Type: types.DefinitionTypeEnum},
				"Name":   {Type: types.DefinitionTypeString},
			},
			expected: true,
		},
		{
			name: "No enum",
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hasEnum(tt.definitions)
			if result != tt.expected {
				t.Errorf("hasEnum() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestHasTemplate(t *testing.T) {
	tests := []struct {
		name        string
		definitions map[string]types.Definition
		expected    bool
	}{
		{
			name: "Has template",
			definitions: map[string]types.Definition{
				"ApiUrl": {Type: types.DefinitionTypeTemplate},
				"Name":   {Type: types.DefinitionTypeString},
			},
			expected: true,
		},
		{
			name: "No template",
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hasTemplate(tt.definitions)
			if result != tt.expected {
				t.Errorf("hasTemplate() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestTryParseDate(t *testing.T) {
	tests := []struct {
		input       string
		shouldParse bool
	}{
		{"2023-01-01T00:00:00Z", true},
		{"2023-12-31T23:59:59Z", true},
		{"invalid-date", false},
		{"", false},
		{"2023-01-01", false}, // RFC3339形式ではない
	}

	for _, tt := range tests {
		_, ok := tryParseDate(tt.input)
		if ok != tt.shouldParse {
			t.Errorf("tryParseDate(%q) parsing success = %v, expected %v", tt.input, ok, tt.shouldParse)
		}
	}
}
