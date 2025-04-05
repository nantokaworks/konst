package template

const defaultGoTemplate = `package {{ .GoPackage }}

{{- if hasDate .Definitions }}
import "time"
{{- end }}

{{- range $name, $def := .Definitions }}
	{{- if eq $def.Type "date" }}
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
