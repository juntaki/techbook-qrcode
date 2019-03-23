package {{.Models}}

{{$ilen := len .Imports}}
{{if gt $ilen 0}}
import (
  {{range .Imports}}"{{.}}"{{end}}
)
{{end}}
{{range .Tables}}
type {{Mapper .Name}} struct {
{{$table := .}}
{{range .ColumnsSeq}}{{$col := $table.GetColumn .}}
  {{Mapper $col.Name}}  {{Type $col}} `db:"{{$col.Name}}"`{{end}}
}
{{end}}
