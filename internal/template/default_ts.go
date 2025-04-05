package template

const defaultTSTemplate = `{{- range $name, $def := .Definitions -}}
export const {{ $name }} = {{ formatTS $def }};
{{ end }}`
