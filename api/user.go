package api

import (
	"errors"
	"fmt"
	db "main/db/sqlc"
	"main/util"
	"main/views"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) signup(ctx *gin.Context) {
	c := views.Signup()
	err := views.Layout(c, "Postings", views.SIGNUP_TAB, false).Render(ctx, ctx.Writer)
	if err != nil {
		http.Error(ctx.Writer, "Error rendering home template", http.StatusInternalServerError)
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	email := ctx.Request.Form.Get("email")
	password := ctx.Request.Form.Get("password")
	username := ctx.Request.Form.Get("username")

	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:       username,
		HashedPassword: hashedPassword,
		Email:          email,
		Role:           "USER",
	}

	user, err := server.store.CreateUser(ctx, arg)
	fmt.Printf("User name is: %s", user.Username)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c := views.Index()
	err = views.Layout(c, "Postings", views.HOME_TAB, false).Render(ctx, ctx.Writer)
	if err != nil {
		http.Error(ctx.Writer, "Error rendering home template", http.StatusInternalServerError)
	}
}

func (server *Server) login(ctx *gin.Context) {
	email := ctx.Request.Form.Get("email")
	password := ctx.Request.Form.Get("password")

	user, err := server.store.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, _, err := server.tokenMaker.CreateToken(
		user.Username,
		user.Role,
		server.config.AccessTokenDuration,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.SetCookie(server.config.AccessTokenCookieName, accessToken, -1, "/", "", true, true)
	c := views.Index()
	err = views.Layout(c, "Postings", views.HOME_TAB, false).Render(ctx, ctx.Writer)
	if err != nil {
		http.Error(ctx.Writer, "Error rendering home template", http.StatusInternalServerError)
	}
}
