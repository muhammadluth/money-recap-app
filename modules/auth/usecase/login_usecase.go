package usecase

import (
	"money-recap-app/app/log"
	"money-recap-app/modules/auth"

	"github.com/gofiber/fiber/v2"
)

type LoginUsecase struct {
}

func NewLoginUsecase() auth.ILoginUsecase {
	return &LoginUsecase{}
}

func (u *LoginUsecase) Login(traceId string, ctx *fiber.Ctx) error {
	log.Event(traceId, "HALO")
	return ctx.JSON(fiber.Map{
		"message": "HALO LOGIN",
	})
}
