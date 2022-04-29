package main

import (
	"html/template"
	"log"
	"os"
)

func imageListToMarkdown(images []Image) string {
	temp := template.Must(template.ParseFiles("index.md.tmpl"))
	err := temp.Execute(os.Stdout, images)
	if err != nil {
		log.Fatalln(err)
	}
	return ""
}
