package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ImageIndex(ctx *gin.Context) {
	ctx.Redirect(http.StatusTemporaryRedirect, "/templates/image_index.html")
}
