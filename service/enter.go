package service

type ServiceGroup struct {
	EsService
	BaseService
	JwtService
	GaodeService
}

var ServiceGroupApp = new(ServiceGroup)
