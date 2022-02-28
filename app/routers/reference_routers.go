package routers

import (
	"money-recap-app/app/utils"
	"money-recap-app/model/constant"
	"money-recap-app/modules/reference"

	"github.com/gofiber/fiber/v2"
)

type ReferenceRouters struct {
	iBankUsecase reference.IBankUsecase
}

func NewReferenceRouters(fiberApp *fiber.App, fiberRouter fiber.Router, iBankUsecase reference.IBankUsecase) {
	routers := &ReferenceRouters{iBankUsecase}

	moduleRoot := fiberRouter.Group("reference")

	// Bank
	moduleRoot.Post("/bank", routers.CreateBank)

	// Dropdown
	moduleRoot.Get("/dropdown/bank", routers.GetDropdownBank)
}

// Bank
func (r *ReferenceRouters) CreateBank(ctx *fiber.Ctx) error {
	traceId, _ := ctx.Locals(constant.CONTEXT_LOCALS_KEY_TRACE_ID).(string)
	if traceId == "" {
		traceId = utils.CreateTraceID()
	}
	return r.iBankUsecase.CreateBank(traceId, ctx)
}

// Dropdown
func (r *ReferenceRouters) GetDropdownBank(ctx *fiber.Ctx) error {
	traceId, _ := ctx.Locals(constant.CONTEXT_LOCALS_KEY_TRACE_ID).(string)
	if traceId == "" {
		traceId = utils.CreateTraceID()
	}
	return r.iBankUsecase.GetDropdownBank(traceId, ctx)
}
