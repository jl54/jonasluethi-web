package responders

import (
	"html/template"
	"log"
	"net/http"
)

func ResponWithHtml(w http.ResponseWriter, layoutFile, templateFile string, data any) {
	t := template.New("layout").Funcs(template.FuncMap{"mod": func(i, j int) bool { return i%j == 0 }})
	t, err := t.ParseFiles(layoutFile, templateFile)

	if err != nil {
		log.Fatalf("failed to parse files: %v", err)
	}
	
	t, err = t.ParseGlob("web/template/blocks/*.html")

	if err != nil {
		log.Fatalf("failed to parse glob: %v", err)
	}

	err = t.ExecuteTemplate(w, "layout", data)

	if err != nil {
		log.Fatalf("failed to parse template: %v", err)
	}
}
