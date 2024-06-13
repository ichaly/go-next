# dbinfo:{{if .Type}}{{ .Type }}{{else}}postgres{{end}},{{- .Version }},{{- .Schema }}

{{ define "schema_directive"}}
{{- if and (ne .Schema "public") (ne .Schema "")}} @schema(name: {{ .Schema }}){{end}}
{{- end}}

{{- define "relation_directive"}}
{{- if (ne .FKeyTable "")}} @relation(type: {{ .FKeyTable }}
{{- if (ne .FKeyCol "")}}, field: {{ .FKeyCol }}{{end -}}
{{- if and (ne .FKeySchema "public") (ne .FKeySchema "")}}, schema: {{ .FKeySchema }}{{end -}})
{{- end}}
{{- end}}

{{- define "function_directive"}}
{{- " @function" }}
{{- if (ne .Type "")}}(return_type: {{ .Type }}){{end}}
{{- end}}

{{- define "column_type"}}
{{- $var := .Type|dbtype }}
{{- $type := (index $var 0)|pascal }}
{{- if .Array}}[{{ $type }}]{{else}}{{ $type }}{{end}}
{{- if .NotNull}}!{{end}}
{{- "\t" }}
{{- if ne (index $var 1) ""}} @type(args: {{ (index $var 1) | printf "%q" }}){{end}}
{{- template "relation_directive" .}}
{{- end}}

{{- define "column"}}
{{ "\t" }}
{{- .Name }}:
{{- "\t"}}
{{- template "column_type" .}}
{{- if .PrimaryKey}} @id{{end}}
{{- if .UniqueKey}} @unique{{end}}
{{- if .FullText}} @search{{end}}
{{- if .Blocked}} @blocked{{end}}
{{- end}}

{{- define "func_args"}}
{{ "\t" }}
{{- .Name }}:
{{- "\t"}}
{{- $var := .Type|dbtype }}
{{- (index $var 0)|pascal }}
{{- if .Array}}[]{{end}}
{{- "\t"}}
{{- if ne (index $var 1) ""}} @type_args({{ (index $var 1) }}){{end}}
{{- end -}}

{{range .Tables -}}
type {{.Name}}
{{- template "schema_directive" .}} {
{{- range .Columns}}{{template "column" .}}{{end}}
}

{{end -}}

{{range .Functions -}}
type {{.Name}}
{{- template "schema_directive" .}}
{{- template "function_directive" .}} {
{{- range .Inputs}}{{template "func_args" .}}{{"\t"}}@input{{end}}
{{- range .Outputs}}{{template "func_args" .}}{{"\t"}}@output{{end}}
}

{{end -}}