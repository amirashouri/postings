package api

import (
	"main/views"
	"net/http"
)

func (server *Server) createUser(w http.ResponseWriter, r *http.Request) {
	c := views.Index()
	err := views.Layout(c, "Postings", views.HOME_TAB, false).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering home template", http.StatusInternalServerError)
	}
}
