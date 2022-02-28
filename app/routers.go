package app

import (
	"money-recap-app/app/routers"
	"money-recap-app/model/dto"
	"money-recap-app/modules/auth"
	"money-recap-app/modules/reference"

	"github.com/gofiber/fiber/v2"
)

type SetupRouters struct {
	propertiesService dto.PropertiesService
	iLoginUsecase     auth.ILoginUsecase
	iBankUsecase      reference.IBankUsecase
}

func NewSetupRouters(propertiesService dto.PropertiesService, iLoginUsecase auth.ILoginUsecase,
	iBankUsecase reference.IBankUsecase) SetupRouters {
	return SetupRouters{propertiesService, iLoginUsecase, iBankUsecase}
}

func (r *SetupRouters) InitRouters(fiberApp *fiber.App) {
	v1Group := fiberApp.Group("v1")

	// AUTH - Login
	routers.NewAuthRouters(fiberApp, v1Group, r.iLoginUsecase)

	// REFERENCE - Bank

	routers.NewReferenceRouters(fiberApp, v1Group, r.iBankUsecase)
}
