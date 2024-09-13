"""
A cursor is an encoded string use for pagination
"""
scalar Cursor
"""
The `DateTime` scalar type represents a DateTime. The DateTime is serialized as an RFC 3339 quoted string
"""
scalar DateTime

{{- range $key,$obj := . }}
{{- if $obj.Description }}
"""
{{ $obj.Description }}
"""
{{- end }}
type {{ $key }} {
{{- range $f := $obj.Fields }}
    {{- if $f.Description }}
    """
    {{ $f.Description }}
    """
    {{- end }}
    {{ $f.Name }}: {{ $f.Type }}
{{- end }}
}
{{ end }}