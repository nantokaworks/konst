package template

const defaultTSTemplate = `{{- range $name, $def := .Definitions -}}
{{- if eq $def.Type "enum" }}
// {{ $name }} enum values
export const {{ $name }} = {
	{{- range $value := $def.Values }}
	{{ toTitle $value }}: "{{ $value }}",
	{{- end }}
} as const;

export type {{ $name }}Type = typeof {{ $name }}[keyof typeof {{ $name }}];

// Type guard for {{ $name }}
export function isValid{{ $name }}(value: string): value is {{ $name }}Type {
	return Object.values({{ $name }}).includes(value as {{ $name }}Type);
}

// Parser for {{ $name }} with exception
export function parse{{ $name }}(value: string): {{ $name }}Type {
	if (isValid{{ $name }}(value)) {
		return value;
	}
	throw new Error("Invalid {{ $name }}: " + value);
}

// Safe parser for {{ $name }} returning undefined on error
export function parse{{ $name }}Safe(value: string): {{ $name }}Type | undefined {
	return isValid{{ $name }}(value) ? value : undefined;
}

// Get all {{ $name }} values
export function getAll{{ $name }}Values(): {{ $name }}Type[] {
	return Object.values({{ $name }});
}

{{- if $def.Default }}
// Get default {{ $name }} value
export function getDefault{{ $name }}(): {{ $name }}Type {
	return {{ $name }}.{{ toTitle $def.Default }};
}
{{- end }}

{{- else }}
export const {{ $name }} = {{ formatTS $def }};
{{- end }}
{{ end }}`
