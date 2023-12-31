package handlers

import (
	"net/http"

	"github.com/enirox001/shortit/components"
)

func Login(w http.ResponseWriter, r *http.Request) {
	component := components.Login()
	component.Render(r.Context(), w)
}
