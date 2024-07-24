scalar DateTime

{{- range $key,$table := . }}
{{- if $table.Description }}
"""
{{ $table.Description }}
"""
{{- end }}
type {{ $key }} {
{{- range $table.Columns }}
    {{- if .Description }}
    """
    {{ .Description }}
    """
    {{- end }}
    {{ .Name }}: {{ .Type }}
{{- end }}
}
{{ end }}
type Query {
{{- range $key,$table := . }}
    {{ toLowerCamel $key }}: {{ $key }}
{{- end }}
}