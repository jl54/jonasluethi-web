package handlers

import (
	"net/http"

	"github.com/jl54/jonasluethi-web/internal/responders"
	"github.com/jl54/jonasluethi-web/internal/store"
)

type PageHandler struct {}

func (handler PageHandler) Handle(w http.ResponseWriter, r *http.Request) {
	responders.RespondWithHtml(w, store.GetPage(r))	
}
