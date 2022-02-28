package main

import (
	"fmt"
	"money-recap-app/app"
	"money-recap-app/app/log"

	// AUTH - Login
	auth_usecase "money-recap-app/modules/auth/usecase"

	// REFERENCE - Bank
	reference_mapper "money-recap-app/modules/reference/mapper"
	reference_repository "money-recap-app/modules/reference/repository"
	reference_usecase "money-recap-app/modules/reference/usecase"

	"github.com/gofiber/fiber/v2"
)

func RunApplication() {
	fmt.Println("Init Configuration...")
	propertiesService := getPropertiesService()

	log.SetupLogging(propertiesService.AppModeDebug)

	dbPostgres := dbPostgresConnect(propertiesService.AppModeDebug, propertiesService.DbPostgres)

	generateConfigFirebaseAdminSDK(dbPostgres, propertiesService.Firebase)

	fiberApp := fiber.New(serviceConfig())

	// AUTH - Login
	iLoginUsecase := auth_usecase.NewLoginUsecase()

	// REFERENCE - Bank
	iBankMapper := reference_mapper.NewBankMapper()
	iBankRepository := reference_repository.NewBankRepository(dbPostgres)
	iBankUsecase := reference_usecase.NewBankUsecase(iBankMapper, iBankRepository)

	iSetupRouters := app.NewSetupRouters(propertiesService, iLoginUsecase, iBankUsecase)
	iSetupRouters.InitRouters(fiberApp)

	iSetupServer := app.NewSetupServer(propertiesService)
	iSetupServer.InitServer(fiberApp)
}
