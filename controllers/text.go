package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/guoruibiao/webcopy/library/localcache"
	"net/http"
)

var CONTENT string

func TextIndex(ctx *gin.Context) {
	ctx.Redirect(http.StatusTemporaryRedirect, "/templates/text_index.html")
}

func TextPage(ctx *gin.Context) {
	data := make(map[string]int)
	mapper := map[string]interface{}{
		"data": data,
	}

	ctx.JSON(http.StatusOK, mapper)
}

func TextGet(ctx *gin.Context) {
	data := make(map[string]interface{})
	data["content"] = localcache.LOCALCACHE

	fmt.Println("----------------")
	fmt.Println(data)
	fmt.Println("----------------")
	mapper := map[string]interface{}{
		"data": data,
	}
	ctx.JSON(http.StatusOK, mapper)
}

func TextUpdate(ctx *gin.Context) {
	content, _ := ctx.GetPostForm("content")

	data := make(map[string]interface{})
	data["content"] = content
	localcache.LOCALCACHE = content

	if content != "" {
		data["status"] = "1"
	}else{
		data["status"] = "0"
	}

	mapper := map[string]interface{}{
		"data": data,
	}
	ctx.JSON(http.StatusOK, mapper)
}
