package auth

import "github.com/gofiber/fiber/v2"

type ILoginUsecase interface {
	Login(traceId string, ctx *fiber.Ctx) error
}
