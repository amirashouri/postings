package api

import (
	db "main/db/sqlc"
	"main/views"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) indexHandler(ctx *gin.Context) {
	c := views.Index()
	err := views.Layout(c, "Postings", views.INDEX_TAB, false).Render(ctx, ctx.Writer)
	if err != nil {
		http.Error(ctx.Writer, "Error rendering home template", http.StatusInternalServerError)
	}
}

func (server *Server) homeHandler(ctx *gin.Context) {
	params := db.GetPostsParams{Limit: 10, Offset: 0}
	posts, err := server.store.GetPosts(ctx, params)
	if err != nil {
		http.Error(ctx.Writer, "Error fetching user posts", http.StatusInternalServerError)
	}
	c := views.Home(posts, true)
	err = views.Layout(c, "Postings", views.HOME_TAB, true).Render(ctx, ctx.Writer)
	if err != nil {
		http.Error(ctx.Writer, "Error rendering home template", http.StatusInternalServerError)
	}
}
