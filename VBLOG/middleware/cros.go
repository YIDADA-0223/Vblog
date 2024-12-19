package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CROS(ctx *gin.Context) {
	h := ctx.Writer.Header()
	if ctx.Request.Method == http.MethodOptions {
		// 复杂请求: preflight
		h["Access-Control-Allow-Headers"] = []string{"Origin,Content-Length,Content-Type"}
		h["Access-Control-Allow-Methods"] = []string{"GET,POST,PUT,PATCH,DELETE,HEAD,OPTIONS"}
		h["Access-Control-Max-Age"] = []string{"43200"}
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	} else {
		// 简单请求
		h["Access-Control-Allow-Origin"] = []string{"*"}
	}
	ctx.Next()
}

/*
var request = new XMLHttpRequest();
request.onreadystatechange = function () {
	console.log(request.status)
	console.log(request.responseText)
}
request.open('GET','http://127.0.0.1:8080/vblog/api/v1/blogs');
request.send();
*/
