package template

const defaultGoTemplate = `package {{ .GoPackage }}
{{- $needsStrings := or (hasEnum .Definitions) (hasTemplate .Definitions) }}
{{- $needsTime := hasDate .Definitions }}
{{- if or $needsStrings $needsTime }}
import (
{{- if or $needsStrings $needsTime }}
	"errors"
{{- end }}
{{- if $needsStrings }}
	"strings"
{{- end }}
{{- if $needsTime }}
	"time"
{{- end }}
)
{{- end }}

{{- range $name, $def := .Definitions }}
	{{- if eq $def.Type "template" }}
// {{ $name }} template string
const {{ $name }}Template = {{ printf "%q" $def.Template }}

// Build{{ $name }} builds the template string with provided parameters
func Build{{ $name }}({{- range $i, $param := $def.Parameters }}{{if $i}}, {{end}}{{ toCamel $param }} string{{- end }}) string {
	result := {{ $name }}Template
	{{- range $param := $def.Parameters }}
	result = strings.ReplaceAll(result, "%{{ $param }}%", {{ toCamel $param }})
	{{- end }}
	return result
}

	{{- else if eq $def.Type "enum" }}
// {{ $name }} enum values
type {{ $name }} string

const (
	{{- range $i, $value := $def.Values }}
	{{ $name }}{{ toTitle $value }} {{ $name }} = "{{ $value }}"
	{{- end }}
)

// IsValid{{ $name }} validates if the given string is a valid {{ $name }}
func IsValid{{ $name }}(value string) bool {
	switch value {
	{{- range $value := $def.Values }}
	case "{{ $value }}":
		return true
	{{- end }}
	default:
		return false
	}
}

// Parse{{ $name }} parses a string to {{ $name }} with error handling
func Parse{{ $name }}(value string) ({{ $name }}, error) {
	if IsValid{{ $name }}(value) {
		return {{ $name }}(value), nil
	}
	return "", errors.New("invalid {{ $name }}: " + value)
}

// GetAll{{ $name }}Values returns all valid {{ $name }} values
func GetAll{{ $name }}Values() []{{ $name }} {
	return []{{ $name }}{
		{{- range $value := $def.Values }}
		{{ $name }}{{ toTitle $value }},
		{{- end }}
	}
}

{{- if $def.Default }}
// GetDefault{{ $name }} returns the default {{ $name }} value
func GetDefault{{ $name }}() {{ $name }} {
	return {{ $name }}{{ toTitle $def.Default }}
}
{{- end }}

	{{- else if eq $def.Type "date" }}
		{{- if eq $def.GoMode "string" }}
const {{ $name }} = {{ formatConstValue $def }}
		{{- else }}
var {{ $name }} = {{ formatConstValue $def }}
		{{- end }}
	{{- else if (contains (asString $def.Type) "[]") }}
var {{ $name }} = {{ formatConstValue $def }}
	{{- else if and (ne $def.Type "template") (ne $def.Type "enum") }}
const {{ $name }} = {{ formatConstValue $def }}
	{{- end }}
{{- end }}`
