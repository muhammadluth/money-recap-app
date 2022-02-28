package app

import (
	"fmt"
	"money-recap-app/model/dto"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type SetupServer struct {
	propertiesService dto.PropertiesService
}

func NewSetupServer(propertiesService dto.PropertiesService) SetupServer {
	return SetupServer{propertiesService}
}

func (s *SetupServer) InitServer(fiberApp *fiber.App) {
	// Config Service
	fiberApp.Use(etag.New())
	fiberApp.Use(compress.New())
	fiberApp.Use(requestid.New())
	fiberApp.Use(recover.New())
	fiberApp.Use(cors.New(cors.Config{
		Next:         cors.ConfigDefault.Next,
		AllowOrigins: s.propertiesService.Cors.AllowOrigins,
		AllowMethods: fmt.Sprintf("%s, %s, %s, %s, %s, %s",
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodPut,
			fiber.MethodPatch,
			fiber.MethodDelete,
			fiber.MethodOptions),
		AllowHeaders:     "*",
		AllowCredentials: true,
		ExposeHeaders:    "*",
		MaxAge:           cors.ConfigDefault.MaxAge,
	}))

	// HEALTH CHECK
	fiberApp.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"message": "Hello 👋, Welcome to Money Recap App!"})
	})

	svcPort := fmt.Sprint(s.propertiesService.Service.ServicePort)
	fmt.Printf("Listening on port : %s\n", svcPort)
	fmt.Printf("Ready to serve\n")
	fiberApp.Listen(fmt.Sprintf(":%s", svcPort))
}
