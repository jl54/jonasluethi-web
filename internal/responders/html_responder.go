package responders

import (
	"html/template"
	"log"
	"net/http"

	"github.com/jl54/jonasluethi-web/internal/models"
)

func RespondWithHtml(w http.ResponseWriter, data *models.Page) {
	t := template.New("layout").Funcs(template.FuncMap{
		"mod": func(i, j int) bool { return i%j == 0 },
	})

	t, err := t.ParseFiles(data.Layout, "web/template/content.html")

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
