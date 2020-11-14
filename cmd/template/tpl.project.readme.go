package template

func init() {
	ProjectTpl["README.MD"] = `
# {{ .ProjectName }}
`
}
