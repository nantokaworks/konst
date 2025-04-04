package template

// 内蔵テンプレート（Go用）
const defaultGoTemplate = `package {{ .PackageName }}

{{- if hasDate .Definitions }}
import "time"
{{- end }}

{{- range $name, $def := .Definitions }}
  {{- if eq $def.Descriptor "const" }}
// {{ $name }} 定数
{{- with $def.Content.ConstContent }}
    {{- if eq .Type "date" }}
var {{ $name }} = {{ formatConstValue $def.Content }}
    {{- else }}
const {{ $name }} = {{ formatConstValue $def.Content }}
    {{- end }}
{{- else }}
const {{ $name }} = {{ formatConstValue $def.Content }}
{{- end }}

  
  {{- end }}

  {{- if eq $def.Descriptor "enum" }}
type {{ $name }} int

const (
  {{- $keys := sortedKeys $def.Content.EnumContent.Values }}
  {{- range $i, $key := $keys }}
    {{- if eq (printf "%v" $i) "0" }}
{{ indent 1 (print $name (title $key) " " $name " = iota") }}
    {{- else }}
{{ indent 1 (print $name (title $key)) }}
    {{- end }}
  {{- end }}
)
  
func (e {{ $name }}) String() string {
{{ indent 1 (print "switch e {") }}
  {{- range $i, $key := $keys }}
    {{- if ne (printf "%v" $i) "0" }}
{{ indent 1 (print "case " $name (title $key) ":") }}
{{ indent 2 (print "return " (printf "%q" (index $def.Content.EnumContent.Values $key))) }}
    {{- end }}
  {{- end }}
{{ indent 1 (print "default:") }}
{{ indent 2 (print "return " (printf "%q" (asString (index $def.Content.EnumContent.Values (index $keys 0))))) }}
{{ indent 1 "}" }}
}
  
{{- end }}

  {{- if eq $def.Descriptor "object" }}
// {{ $name }} 構造体
type {{ $name }} struct {
{{- range $field, $info := $def.Content.ObjectContent.Fields }}
{{ indent 1 (print $field " " $info.Type) }}
{{- end }}
}
  
  {{- end }}
{{- end }}`
