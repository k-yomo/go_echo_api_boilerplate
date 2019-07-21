{{- $short := (shortname .Type.Name) -}}
// {{ .Name }} returns the {{ .RefType.Name }} associated with the {{ .Type.Name }}'s {{ .Field.Name }} ({{ .Field.Col.ColumnName }}).
//
// Generated from foreign key '{{ .ForeignKey.ForeignKeyName }}'.
func ({{ $short }} *{{ .Type.Name }}) {{ .Name }}(ctx context.Context, db Executor) (*{{ .RefType.Name }}, error) {
	return {{ .RefType.Name }}By{{ .RefField.Name }}(ctx, db, {{ convext $short .Field .RefField }})
}

