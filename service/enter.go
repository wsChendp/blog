package service

type ServiceGroup struct {
	EsService
	BaseService
	JwtService
	GaodeService
	QQService
	UserService
}

var ServiceGroupApp = new(ServiceGroup)
