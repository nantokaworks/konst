package utils

import (
	"path/filepath"
	"strings"
	"unicode"
)

// ToSnakeCase converts a string to snake_case
func ToSnakeCase(s string) string {
	if s == "" {
		return s
	}

	var result []rune
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			prev := rune(s[i-1])
			// Add underscore before uppercase if previous is lowercase or digit
			if unicode.IsLower(prev) || unicode.IsDigit(prev) {
				result = append(result, '_')
			}
		}
		// Replace hyphens with underscores
		if r == '-' {
			result = append(result, '_')
		} else {
			result = append(result, unicode.ToLower(r))
		}
	}
	return string(result)
}

// ToKebabCase converts a string to kebab-case
func ToKebabCase(s string) string {
	if s == "" {
		return s
	}

	var result []rune
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			prev := rune(s[i-1])
			// Add hyphen before uppercase if previous is lowercase or digit
			if unicode.IsLower(prev) || unicode.IsDigit(prev) {
				result = append(result, '-')
			}
		}
		// Replace underscores with hyphens
		if r == '_' {
			result = append(result, '-')
		} else {
			result = append(result, unicode.ToLower(r))
		}
	}
	return string(result)
}

// ToCamelCase converts a string to camelCase
func ToCamelCase(s string) string {
	if s == "" {
		return s
	}

	// Split by underscores and hyphens
	parts := strings.FieldsFunc(s, func(r rune) bool {
		return r == '_' || r == '-'
	})

	if len(parts) == 0 {
		return s
	}

	// First part stays lowercase
	result := strings.ToLower(parts[0])

	// Capitalize first letter of remaining parts
	for i := 1; i < len(parts); i++ {
		if len(parts[i]) > 0 {
			result += strings.ToUpper(string(parts[i][0])) + strings.ToLower(parts[i][1:])
		}
	}

	return result
}

// ConvertFileName converts a file name to the specified naming style
func ConvertFileName(fileName string, namingStyle string, isTS bool) string {
	// Extract base name without extension
	base := fileName
	if idx := strings.LastIndex(fileName, "."); idx != -1 {
		base = fileName[:idx]
	}

	// Apply default styles if no style specified
	if namingStyle == "" {
		if isTS {
			namingStyle = "kebab"
		} else {
			namingStyle = "snake"
		}
	}

	// Convert based on style
	switch namingStyle {
	case "kebab":
		return ToKebabCase(base)
	case "camel":
		return ToCamelCase(base)
	case "snake":
		return ToSnakeCase(base)
	default:
		// If invalid style, return original base
		return base
	}
}

// ConvertPath converts a file path including directories to the specified naming style
func ConvertPath(path string, namingStyle string, isTS bool) string {
	// Split path into directory components
	parts := strings.Split(path, string(filepath.Separator))
	
	// Convert each part
	for i := range parts {
		parts[i] = ConvertFileName(parts[i], namingStyle, isTS)
	}
	
	// Rejoin with proper separator
	return filepath.Join(parts...)
}