package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/atomdata/go-man-page/data"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := getMainTemplate()
		if t != nil {
			ad := data.GetManPage()
			err := t.Execute(w, ad)
			if err != nil {
				log.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
	http.Handle("/css/", http.FileServer(http.Dir("public")))
	// js folder inside public for your analytics
	http.Handle("/js/", http.FileServer(http.Dir("public")))

	http.ListenAndServe(":8080", nil)
}

func getMainTemplate() *template.Template {
	funcMap := template.FuncMap{
		"ToUpper":        strings.ToUpper,
		"ToLower":        strings.ToLower,
		"GetIndentClass": data.GetIndentClass,
	}
	result := template.New("man").Funcs(funcMap)
	const basePath = "views"
	template.Must(result.ParseGlob(basePath + "/*.html"))
	template := result.Lookup("page.html")
	return template
}
