package service

type ServiceGroup struct {
	EsService
	BaseService
	JwtService
	GaodeService
	QQService
	UserService
	ImageService
	ArticleService
	CommentService
}

var ServiceGroupApp = new(ServiceGroup)
