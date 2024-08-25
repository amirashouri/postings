package api

import (
	db "main/db/sqlc"
	"main/views"
	"main/views/model"
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
	var postItems []model.PostItem
	for i := 0; i < len(posts); i++ {
		likes, _ := server.store.GetLikes(ctx, posts[i].ID)
		isLiked := contains(likes, posts[i].ID)
		postItem := model.PostItem{Post: posts[i], LikesCount: len(likes), IsLiked: isLiked}
		postItems = append(postItems, postItem)
	}
	c := views.Home(postItems, true)
	err = views.Layout(c, "Postings", views.HOME_TAB, true).Render(ctx, ctx.Writer)
	if err != nil {
		http.Error(ctx.Writer, "Error rendering home template", http.StatusInternalServerError)
	}
}

func contains(s []db.Like, id int64) bool {
	for _, a := range s {
		if a.PostID == id {
			return true
		}
	}
	return false
}
