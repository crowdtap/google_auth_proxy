package main

import (
	"html/template"
	"log"
)

func getTemplates() *template.Template {
	t, err := template.ParseFiles("templates/css.html", "templates/sign_in.html", "templates/error.html")

	if err != nil {
		log.Fatalf("failed parsing template %s", err.Error())
	}

	return t
}
