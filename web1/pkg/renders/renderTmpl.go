package renders

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplates(w http.ResponseWriter, str string) {
	parseTmpl, _ := template.ParseFiles("./templates/" + str)
	err := parseTmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Test
func RenderTmplTest(w http.ResponseWriter, str string) {
	_ = map[string]*template.Template{}

	//
}
