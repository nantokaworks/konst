package template

// 内蔵テンプレート（TypeScript用）
const defaultTSTemplate = `{{- range $name, $def := .Definitions }}
  {{- if eq $def.Descriptor "const" }}
// {{ $name }} 定数
{{- with $def.Content.ConstContent }}
  {{- if eq .Type "date" }}
    {{- if eq .TSMode "string" }}
export const {{ $name }} = {{ formatTS $def.Content }};
    {{- else }}
export const {{ $name }} = new Date({{ formatTS $def.Content }});
    {{- end }}
  {{- else }}
export const {{ $name }} = {{ formatTS $def.Content }};
  {{- end }}
{{- else }}
export const {{ $name }} = {{ formatTS $def.Content }};
{{- end }}

  
  {{- end }}

  {{- if eq $def.Descriptor "enum" }}
// {{ $name }} 列挙型
export const {{ $name }} = {
{{- range $key, $val := $def.Content.EnumContent.Values }}
{{ indent 1 (print $key ": " (printf "%q" $val) ",") }}
{{- end }}
} as const;

export type {{ $name }} = typeof {{ $name }}[keyof typeof {{ $name }}];

  
  {{- end }}

  {{- if eq $def.Descriptor "object" }}
// {{ $name }} オブジェクトインターフェース
export interface {{ $name }} {
{{- range $field, $info := $def.Content.ObjectContent.Fields }}
{{ indent 1 (print $field ": " (convertTSType $info.Type) ";") }}
{{- end }}
}
  
  {{- end }}
{{- end }}`
