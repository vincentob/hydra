package template

func init() {
	ProjectTpl["go.mod"] = `module {{ .ProjectPath }}

go 1.15

require ()
`
}
