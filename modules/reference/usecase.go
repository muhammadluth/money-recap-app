package reference

import "github.com/gofiber/fiber/v2"

type IBankUsecase interface {
	CreateBank(traceId string, ctx *fiber.Ctx) error
	GetDropdownBank(traceId string, ctx *fiber.Ctx) error
}
