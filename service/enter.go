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
}

var ServiceGroupApp = new(ServiceGroup)
