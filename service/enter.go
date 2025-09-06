package service

type ServiceGroup struct {
	EsService
	BaseService
	JwtService
	GaodeService
	QQService
	UserService
	ImageService
}

var ServiceGroupApp = new(ServiceGroup)
