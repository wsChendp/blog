package router

import "github.com/gin-gonic/gin"

type RouterGroup struct {
	BaseRouter
}

var RouterGroupApp = new(RouterGroup)

func (e *RouterGroup) InitEnterRouter(Router *gin.RouterGroup) {

}

func (e *RouterGroup) InitUserRouter(group *gin.RouterGroup, group2 *gin.RouterGroup, group3 *gin.RouterGroup) {

}
