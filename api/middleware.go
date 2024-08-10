package api

import (
	"errors"
	"main/token"
	"main/views"

	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

// AuthMiddleware creates a http middleware for authorization
func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken, err := ctx.Cookie("access-token")
		if err != nil {
			redirect(ctx)
			err = errors.New("authorization header can not be found")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		if len(accessToken) == 0 {
			redirect(ctx)
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			redirect(ctx)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}

func redirect(ctx *gin.Context) {
	c := views.Index(false)
	err := views.Layout(c, "Postings", views.HOME_TAB, false).Render(ctx, ctx.Writer)
	if err != nil {
		http.Error(ctx.Writer, "Error rendering home template", http.StatusInternalServerError)
	}
}
