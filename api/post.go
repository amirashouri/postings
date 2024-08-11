package api

import (
	db "main/db/sqlc"
	"main/token"
	"main/views"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (server *Server) getPosts(ctx *gin.Context) {
	c := views.Posts()
	err := views.Layout(c, "Postings", views.POSTS_TAB, false).Render(ctx, ctx.Writer)
	if err != nil {
		http.Error(ctx.Writer, "Error rendering home template", http.StatusInternalServerError)
	}
}

func (server *Server) createPost(ctx *gin.Context) {
	title := ctx.Request.FormValue("title")
	body := ctx.Request.FormValue("body")

	if title == "" || body == "" {
		// handle error case
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	params := db.CreatePostParams{Title: title, Body: body, UserID: authPayload.UserId, Status: ""}
	_, err := server.store.CreatePost(ctx, params)
	if err != nil {
		log.Err(err).Msg("Server started...")
		http.Error(ctx.Writer, "Error Failed to create a post", http.StatusInternalServerError)
	}
	c := views.Post(title, body)
	err = views.Layout(c, "Postings", views.HOME_TAB, true).Render(ctx, ctx.Writer)
	if err != nil {
		http.Error(ctx.Writer, "Error rendering post template", http.StatusInternalServerError)
	}
}
