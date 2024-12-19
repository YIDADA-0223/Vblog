package response

import (
	"net/http"

	"gitee.com/VBLOG/exception"
	"github.com/gin-gonic/gin"
)

// 成功，怎么把对象返回给HTTP Reponse
func Success(data any, c *gin.Context) {
	//其他逻辑，后期可以做脱敏
	c.JSON(http.StatusOK, data)
}

// 失败，怎么把对象返回给HTTP Reponse
// 统一返回的数据结构：ApiException
func Failed(err error, c *gin.Context) {
	httpCode := http.StatusInternalServerError
	if v, ok := err.(*exception.ApiExcepiton); ok {
		if v.HttpCode != 0 {
			httpCode = v.HttpCode
		}
	} else {
		//非业务异常，直接转换为指定的内部报错异常
		err = exception.ErrServerInternal(err.Error())
	}

	c.JSON(httpCode, err)
	c.Abort()

}
