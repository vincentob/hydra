package template

func init() {
	CommonProjectFiles["README.MD"] = `
# {{ .ProjectName }}
`
}
