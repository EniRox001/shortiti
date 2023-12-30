package handlers

import (
	"net/http"

	"github.com/enirox001/shortit/components"
)

func Home(w http.ResponseWriter, r *http.Request) {
	component := components.DashboardHmoe("starter tem")
	component.Render(r.Context(), w)

}
