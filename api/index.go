package api

import (
	db "main/db/sqlc"
	"main/token"
	"main/views"
	"main/views/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (server *Server) indexHandler(ctx *gin.Context) {
	ok := validateToken(server, ctx)
	if ok {
		ctx.Redirect(http.StatusMovedPermanently, "/home")
		return
	}
	c := views.Index()
	err := views.Layout(c, "Postings", views.INDEX_TAB, false).Render(ctx, ctx.Writer)
	if err != nil {
		http.Error(ctx.Writer, "Error rendering home template", http.StatusInternalServerError)
	}
}

func validateToken(server *Server, ctx *gin.Context) bool {
	accessToken, err := ctx.Cookie("access-token")
	if err != nil {
		return false
	}
	payload, err := server.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return false
	}
	if payload.ExpiredAt.After(time.Now()) {
		return true
	}
	return false
}

func (server *Server) homeHandler(ctx *gin.Context) {
	params := db.GetPostsParams{Limit: 10, Offset: 0}
	posts, err := server.store.GetPosts(ctx, params)
	if err != nil {
		http.Error(ctx.Writer, "Error fetching user posts", http.StatusInternalServerError)
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	var postItems []model.PostItem
	for i := 0; i < len(posts); i++ {

		post := posts[i]
		postItem, err := createPostItem(server, ctx, post.ID, authPayload.UserId)
		if err != nil {
			http.Error(ctx.Writer, "Failed to create a post", http.StatusInternalServerError)
		}
		postItems = append(postItems, postItem)
	}
	c := views.Home(postItems, true)
	err = views.Layout(c, "Postings", views.HOME_TAB, true).Render(ctx, ctx.Writer)
	if err != nil {
		log.Err(err).Msg("Failed to fetch posts")
		http.Error(ctx.Writer, "Error rendering home template", http.StatusInternalServerError)
	}
}
