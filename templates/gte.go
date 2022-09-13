package templates

const GTE = `
{{ define "GTE" }}

var gteC1, gteC2 uint64
{{- range $i := intRange 0 $.limbCount }}
    _, gteC1 = bits.Sub64({{$.z}}[{{$i}}], {{$.x}}[{{$i}}], gteC1)
{{- end}}
{{- range $i := intRange 0 $.limbCount }}
    _, gteC2 = bits.Sub64({{$.z}}[{{$i}}], {{$.y}}[{{$i}}], gteC2)
{{- end}}

if gteC1 != 0 || gteC2 != 0 {
    return errors.New(fmt.Sprintf("input gte modulus"))
}
{{ end }}
`
