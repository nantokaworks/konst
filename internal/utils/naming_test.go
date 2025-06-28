package utils

import (
	"testing"
)

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"already snake case", "hello_world", "hello_world"},
		{"camelCase", "helloWorld", "hello_world"},
		{"PascalCase", "HelloWorld", "hello_world"},
		{"kebab-case", "hello-world", "hello_world"},
		{"mixed case", "someVariableName", "some_variable_name"},
		{"with numbers", "variable123Name", "variable123_name"},
		{"consecutive caps", "XMLParser", "xmlparser"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToSnakeCase(tt.input)
			if result != tt.expected {
				t.Errorf("ToSnakeCase(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToKebabCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"already kebab case", "hello-world", "hello-world"},
		{"camelCase", "helloWorld", "hello-world"},
		{"PascalCase", "HelloWorld", "hello-world"},
		{"snake_case", "hello_world", "hello-world"},
		{"mixed case", "someVariableName", "some-variable-name"},
		{"with numbers", "variable123Name", "variable123-name"},
		{"consecutive caps", "XMLParser", "xmlparser"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToKebabCase(tt.input)
			if result != tt.expected {
				t.Errorf("ToKebabCase(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToCamelCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"already camelCase", "helloWorld", "helloworld"},
		{"snake_case", "hello_world", "helloWorld"},
		{"kebab-case", "hello-world", "helloWorld"},
		{"PascalCase", "HelloWorld", "helloworld"},
		{"mixed case", "some_variable_name", "someVariableName"},
		{"with numbers", "variable_123_name", "variable123Name"},
		{"single word", "hello", "hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToCamelCase(tt.input)
			if result != tt.expected {
				t.Errorf("ToCamelCase(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestConvertFileName(t *testing.T) {
	tests := []struct {
		name        string
		fileName    string
		namingStyle string
		isTS        bool
		expected    string
	}{
		// Default behavior tests
		{"ts default kebab", "test_file", "", true, "test-file"},
		{"go default snake", "test-file", "", false, "test_file"},
		
		// Explicit style tests
		{"explicit kebab", "test_file", "kebab", false, "test-file"},
		{"explicit camel", "test_file", "camel", false, "testFile"},
		{"explicit snake", "test-file", "snake", false, "test_file"},
		
		// File with extension
		{"file with extension", "test_file.json", "kebab", true, "test-file"},
		
		// Invalid style
		{"invalid style", "test_file", "invalid", false, "test_file"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertFileName(tt.fileName, tt.namingStyle, tt.isTS)
			if result != tt.expected {
				t.Errorf("ConvertFileName(%q, %q, %v) = %q, want %q", 
					tt.fileName, tt.namingStyle, tt.isTS, result, tt.expected)
			}
		})
	}
}

func TestConvertPath(t *testing.T) {
	tests := []struct {
		name        string
		path        string
		namingStyle string
		isTS        bool
		expected    string
	}{
		{"single dir kebab", "test_dir", "kebab", true, "test-dir"},
		{"nested dirs kebab", "test_dir/sub_dir", "kebab", true, "test-dir/sub-dir"},
		{"single dir camel", "test-dir", "camel", false, "testDir"},
		{"nested dirs snake", "test-dir/sub-dir", "snake", false, "test_dir/sub_dir"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertPath(tt.path, tt.namingStyle, tt.isTS)
			if result != tt.expected {
				t.Errorf("ConvertPath(%q, %q, %v) = %q, want %q",
					tt.path, tt.namingStyle, tt.isTS, result, tt.expected)
			}
		})
	}
}