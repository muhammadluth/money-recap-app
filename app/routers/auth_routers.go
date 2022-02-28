package routers

import (
	"money-recap-app/app/utils"
	"money-recap-app/model/constant"
	"money-recap-app/modules/auth"

	"github.com/gofiber/fiber/v2"
)

type AuthRouters struct {
	iLoginUsecase auth.ILoginUsecase
}

func NewAuthRouters(fiberApp *fiber.App, fiberRouter fiber.Router, iLoginUsecase auth.ILoginUsecase) {
	routers := &AuthRouters{iLoginUsecase}
	fiberRouter.Post("/auth/login", routers.Login)
}

func (r *AuthRouters) Login(ctx *fiber.Ctx) error {
	traceId, _ := ctx.Locals(constant.CONTEXT_LOCALS_KEY_TRACE_ID).(string)
	if traceId == "" {
		traceId = utils.CreateTraceID()
	}
	return r.iLoginUsecase.Login(traceId, ctx)
}
