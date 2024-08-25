package api

import (
	db "main/db/sqlc"
	"main/token"
	"main/views"
	"main/views/model"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
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
	post, err := server.store.CreatePost(ctx, params)
	if err != nil {
		log.Err(err).Msg("Failed to create post")
		http.Error(ctx.Writer, "Error Failed to create a post", http.StatusInternalServerError)
	}
	postItem := model.PostItem{Post: post, LikesCount: 0, IsLiked: false}
	c := views.Post(postItem)
	err = c.Render(ctx, ctx.Writer)
	if err != nil {
		http.Error(ctx.Writer, "Error rendering post template", http.StatusInternalServerError)
	}
}

func (server *Server) deletePost(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		http.Error(ctx.Writer, "Should provide an id", http.StatusBadRequest)
	}
	err = server.store.DeletePost(ctx, id)
	if err != nil {
		http.Error(ctx.Writer, "Failed to delete the post", http.StatusInternalServerError)
	}
	err = templ.NopComponent.Render(ctx, ctx.Writer)
	if err != nil {
		http.Error(ctx.Writer, "Error rendering home template", http.StatusInternalServerError)
	}
}

func (server *Server) likePost(ctx *gin.Context) {
	// TODO: Fix unlike case and counter update
	idString := ctx.Param("id")
	postId, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		http.Error(ctx.Writer, "Should provide an id", http.StatusBadRequest)
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	likeParam := db.CreateLikeParams{UserID: authPayload.UserId, PostID: postId}
	err = server.store.CreateLike(ctx, likeParam)
	if err != nil {
		http.Error(ctx.Writer, "failed to create like", http.StatusBadRequest)
	}
	c := views.LikeButton(postId, true)
	err = c.Render(ctx, ctx.Writer)
	if err != nil {
		http.Error(ctx.Writer, "Error rendering like button template", http.StatusInternalServerError)
	}
}
