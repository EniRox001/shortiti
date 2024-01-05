package handlers

import (
	"net/http"

	"github.com/enirox001/shortit/components"
	"github.com/enirox001/shortit/data"
	"github.com/enirox001/shortit/util"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	urls := data.FindAll()

	if len(urls) > 0 {
		component := components.DashboardHomeLinks(urls)
		component.Render(r.Context(), w)
		return
	}

	component := components.DashboardHome("starter template")
	component.Render(r.Context(), w)
}

func Shorten(w http.ResponseWriter, r *http.Request) {
	urlString := util.TrimSpace(r.FormValue("link"))

	link, uniqueId := util.GenerateShortLink()

	shortLink := &data.Url{
		ShortLink: link,
		FullLink:  urlString,
		UniqueID:  uniqueId,
		Clicks:    0,
	}

	shortLink.CreateShortUrl()

	http.Redirect(w, r, "/", http.StatusFound)

	// http.Redirect(w, r, "/public/links/"+uniqueId, http.StatusFound)
}

func ShowQRCode(w http.ResponseWriter, r *http.Request) {
	urlId := mux.Vars(r)["urlId"]

	url := data.FindByUniqueId(urlId)

	if url.ID == 0 {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	component := components.DashboardQRCode(*url)
	component.Render(r.Context(), w)

}

func Redirect(w http.ResponseWriter, r *http.Request) {
	urlId := mux.Vars(r)["urlId"]

	url := data.FindByUniqueId(urlId)

	if url.ID == 0 {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	url.Clicks = url.Clicks + 1
	url.UpdateClicks()

	// util.Redirect(url.FullLink)
}

func DeletePublicLink(w http.ResponseWriter, r *http.Request) {
	urlId := mux.Vars(r)["urlId"]

	url := data.FindByUniqueId(urlId)

	if url.ID == 0 {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	url.Delete()

	http.Redirect(w, r, "/", http.StatusFound)
}

func LinkDashboard(w http.ResponseWriter, r *http.Request) {
	urlId := mux.Vars(r)["url"]

	url := data.FindByUniqueId(urlId)

	if url.ID == 0 {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	component := components.LinkDashboard(*url)
	component.Render(r.Context(), w)
}
