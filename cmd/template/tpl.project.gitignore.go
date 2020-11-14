package template

func init() {
	ProjectTpl[".gitignore"] = `# Created by .ignore support plugin (hsz.mobi)
### Go template
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

vendor
.idea/
*.iml
.DS_Store
dist/
{{ .ProjectName }}
playground/

`
}
