package template

func init() {
	CommonProjectFiles["go.mod"] = `module {{ .ProjectPath }}

go 1.15

require ()
`
}
