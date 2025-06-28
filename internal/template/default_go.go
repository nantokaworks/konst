package template

const defaultGoTemplate = `package {{ .GoPackage }}
{{ if hasDate .Definitions }}
import (
	"errors"
	"time"
)
{{ else if hasEnum .Definitions }}
import "errors"
{{ end }}

{{- range $name, $def := .Definitions }}
	{{- if eq $def.Type "enum" }}
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
	{{- else }}
const {{ $name }} = {{ formatConstValue $def }}
	{{- end }}
{{- end }}`
