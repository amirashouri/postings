package api

import (
	"main/views"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) getPosts(ctx *gin.Context) {
	c := views.Posts()
	err := views.Layout(c, "Postings", views.POSTS_TAB, false).Render(ctx, ctx.Writer)
	if err != nil {
		http.Error(ctx.Writer, "Error rendering home template", http.StatusInternalServerError)
	}
}
