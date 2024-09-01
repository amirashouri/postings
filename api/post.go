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

func (server *Server) getPost(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		http.Error(ctx.Writer, "Should provide an id", http.StatusBadRequest)
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	postItem, err := createPostItem(server, ctx, id, authPayload.UserId)
	if err != nil {
		http.Error(ctx.Writer, "failed to create post Item", http.StatusBadRequest)
	}
	c := views.Post(postItem)
	err = c.Render(ctx, ctx.Writer)
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
	postItem, err := createPostItem(server, ctx, post.ID, authPayload.UserId)
	if err != nil {
		http.Error(ctx.Writer, "Failed to create a post item", http.StatusInternalServerError)
	}
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

	param := db.GetLikeByUserParams{UserID: authPayload.UserId, PostID: postId}
	like, notFound := server.store.GetLikeByUser(ctx, param)

	if notFound != nil {
		likeParam := db.CreateLikeParams{UserID: authPayload.UserId, PostID: postId}
		err = server.store.CreateLike(ctx, likeParam)
	} else {
		err = server.store.DeleteLike(ctx, like.ID)
	}
	if err != nil {
		http.Error(ctx.Writer, "failed to create like", http.StatusBadRequest)
	}

	postItem, err := createPostItem(server, ctx, postId, authPayload.UserId)
	if err != nil {
		http.Error(ctx.Writer, "failed to create post Item", http.StatusBadRequest)
	}
	c := views.LikeButton(postItem, notFound != nil)
	err = c.Render(ctx, ctx.Writer)
	if err != nil {
		http.Error(ctx.Writer, "Error rendering like button template", http.StatusInternalServerError)
	}
}

func containsUserIdInLikes(s []db.Like, id int64) (int64, bool) {
	for _, a := range s {
		if a.UserID == id {
			return a.ID, true
		}
	}
	return 0, false
}

func createPostItem(server *Server, ctx *gin.Context, postId int64, userId int64) (model.PostItem, error) {
	post, err := server.store.GetPost(ctx, postId)
	if err != nil {
		return model.PostItem{}, err
	}

	liked, err := server.store.GetLikes(ctx, postId)
	if err != nil {
		return model.PostItem{}, err
	}

	user, err := server.store.GetUser(ctx, post.UserID)
	if err != nil {
		return model.PostItem{}, err
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	_, isLiked := containsUserIdInLikes(liked, authPayload.UserId)

	postItem := model.PostItem{Post: post, LikesCount: len(liked), IsLiked: isLiked, UserName: user.Username, ShowDelete: post.UserID == userId}
	return postItem, nil
}
