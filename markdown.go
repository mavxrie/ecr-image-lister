package main

import (
	"html/template"
	"io"
	"log"
)

const tmpl = `# Docker images

{{ if . -}}
## Last versions

{{- range $value := . }}
- {{.Name}}: {{ if .Versions }}{{ (index .Versions 0) }}{{ else }}{{ (index .RawTags 0) }}{{ end }}
{{- end }}

## All versions

{{- range $value := . }}
- {{.Name}}: {{ if .Versions }}{{ .Versions }}{{ else }}{{ .RawTags }}{{ end }}
{{- end }}
{{- else -}}
No image was found.
{{- end }}
`

func imageListToMarkdown(wr io.Writer, images []Image) {
	temp := template.Must(template.New("").Parse(tmpl))
	err := temp.Execute(wr, images)
	if err != nil {
		log.Fatalln(err)
	}
}
