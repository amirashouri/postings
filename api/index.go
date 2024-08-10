package api

import (
	"main/views"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) indexHandler(ctx *gin.Context) {
	c := views.Index()
	err := views.Layout(c, "Postings", views.HOME_TAB, false).Render(ctx, ctx.Writer)
	if err != nil {
		http.Error(ctx.Writer, "Error rendering home template", http.StatusInternalServerError)
	}
}
