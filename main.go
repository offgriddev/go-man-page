package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/atomdata/go-man-page/data"
)

func main() {
	templates := getMainTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		t := templates.Lookup("page.html")
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
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
	http.Handle("/js/", http.FileServer(http.Dir("public")))

	http.ListenAndServe(":8080", nil)
}

func getMainTemplates() *template.Template {
	funcMap := template.FuncMap{
		"ToUpper": strings.ToUpper,
		"ToLower": strings.ToLower,
	}
	result := template.New("man").Funcs(funcMap)
	const basePath = "views"
	template.Must(result.ParseGlob(basePath + "/*.html"))
	return result
}
