package app

import (
	"money-recap-app/app/log"
	"money-recap-app/app/utils"
	"money-recap-app/model/constant"
	"money-recap-app/model/dto"

	"strings"

	"github.com/gofiber/fiber/v2"
)

type Middleware struct {
	propertiesService dto.PropertiesService
}

func NewMiddleware(propertiesService dto.PropertiesService) Middleware {
	return Middleware{propertiesService}
}

func (m *Middleware) AuthMiddleware() func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var (
			// authorization = ctx.Get("Authorization")
			traceId = utils.CreateTraceID()
		)

		request := strings.ReplaceAll(string(ctx.Request().Body()), "\n", "")
		request = strings.ReplaceAll(request, " ", "")
		log.Message(
			traceId,
			"IN",
			"CLIENT",
			"",
			"GO-FIBER SERVICE",
			"",
			"INCOMING REQUEST",
			string(ctx.Method()),
			string(ctx.OriginalURL()),
			"",
			request,
		)

		// splitAuthorization := strings.Split(authorization, " ")
		// if strings.ToLower(splitAuthorization[0]) != constant.PREFIX_TOKEN {
		// 	err := errors.New("authentication credentials were not provided")
		// 	log.Error(traceId, err)
		// 	return ctx.Status(fiber.StatusUnauthorized).JSON(model.ResponseErrorKeyDetail{
		// 		Detail: strings.Title(err.Error()),
		// 	})
		// } else if len(splitAuthorization) != 2 {
		// 	err := errors.New("invalid token header. no credentials provided")
		// 	log.Error(traceId, err)
		// 	return ctx.Status(fiber.StatusUnauthorized).JSON(model.ResponseErrorKeyDetail{
		// 		Detail: strings.Title(err.Error()),
		// 	})
		// }

		// dataUserRole, dataUserId, err := m.iPermissionUsecase.UserRole(traceId, splitAuthorization[1])
		// if err != nil {
		// 	return ctx.Status(fiber.StatusUnauthorized).JSON(model.ResponseErrorKeyDetail{
		// 		Detail: "Invalid Token",
		// 	})
		// }

		ctx.Set("X-XSS-Protection", "1; mode=block")
		ctx.Set("X-Content-Type-Options", "nosniff")
		ctx.Set("X-Download-Options", "noopen")
		ctx.Set("Strict-Transport-Security", "max-age=5184000")
		ctx.Set("X-Frame-Options", "SAMEORIGIN")
		ctx.Set("X-DNS-Prefetch-Control", "off")

		// ctx.Locals(constant.CONTEXT_LOCALS_KEY_PERMISSION_USER_ROLE, dataUserRole)
		// ctx.Locals(constant.CONTEXT_LOCALS_KEY_USER_ID, dataUserId)
		ctx.Locals(constant.CONTEXT_LOCALS_KEY_TRACE_ID, traceId)

		ctx.Next()

		log.Message(
			traceId,
			"OUT",
			"CLIENT",
			"",
			"GO-FIBER SERVICE",
			"",
			"OUTGOING RESPONSE",
			string(ctx.Method()),
			string(ctx.OriginalURL()),
			"",
			string(ctx.Response().Body()),
		)
		return nil
	}

}
