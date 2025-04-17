package router

import "github.com/gin-gonic/gin"

type RouterGroup struct{
	BaseRouter
}

var RouterGroupApp = new(RouterGroup)

func (e *RouterGroup) InitEnterRouter(Router *gin.RouterGroup) {
	
}
