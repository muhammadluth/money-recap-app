package usecase

import (
	"fmt"
	"money-recap-app/app/log"
	"money-recap-app/model/dto"
	"money-recap-app/modules/reference"

	"github.com/gofiber/fiber/v2"
)

type BankUsecase struct {
	iBankMapper     reference.IBankMapper
	iBankRepository reference.IBankRepository
}

func NewBankUsecase(iBankMapper reference.IBankMapper, iBankRepository reference.IBankRepository) reference.IBankUsecase {
	return &BankUsecase{iBankMapper, iBankRepository}
}

func (u *BankUsecase) CreateBank(traceId string, ctx *fiber.Ctx) error {
	request := new(dto.RequestCreateBank)

	if err := ctx.BodyParser(request); err != nil {
		log.Error(traceId, err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ResponseHttp{
			Message: err.Error(),
		})
	}

	bankData := u.iBankMapper.ToInsertBank(*request)
	if err := u.iBankRepository.CreateBankDB(traceId, bankData); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ResponseHttp{
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(dto.ResponseHttp{
		Message: fmt.Sprintf("Successfully Create Reference Data %s", request.BankName),
	})
}

func (u *BankUsecase) GetDropdownBank(traceId string, ctx *fiber.Ctx) error {
	_, bankData, err := u.iBankRepository.GetListBankDB(traceId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ResponseHttp{
			Message: err.Error(),
		})
	}
	response := u.iBankMapper.ToDropdownBank(bankData)
	return ctx.JSON(response)
}
